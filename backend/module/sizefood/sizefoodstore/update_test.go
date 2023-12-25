package sizefoodstore

import (
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_UpdateSizeFood(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	foodId := "Food001"
	sizeId := "Size001"
	name := mock.Anything
	cost := 0
	price := 0
	mockData := sizefoodmodel.SizeFoodUpdate{
		Name:  &name,
		Cost:  &cost,
		Price: &price,
	}
	expectedQuery := "UPDATE `SizeFood` SET `name`=?,`cost`=?,`price`=? WHERE foodId = ? and sizeId = ?"
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		foodId string
		sizeId string
		data   *sizefoodmodel.SizeFoodUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update size food successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:    context.Background(),
				foodId: foodId,
				sizeId: sizeId,
				data:   &mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.Name, mockData.Cost, mockData.Price, foodId, sizeId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update size food failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "Food001",
				sizeId: "Size001",
				data:   &mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.Name, mockData.Cost, mockData.Price, foodId, sizeId).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}
			tt.mock()

			err := s.UpdateSizeFood(tt.args.ctx, tt.args.foodId, tt.args.sizeId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
