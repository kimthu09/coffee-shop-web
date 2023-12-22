package productstore

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_UpdateTopping(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	id := "123"
	name := "UpdatedToppingName"
	description := "UpdatedDescription"
	cookingGuide := "UpdatedCookingGuide"
	cost := 10
	price := 8
	updateData := &productmodel.ToppingUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &name,
			Description:  &description,
			CookingGuide: &cookingGuide,
		},
		Cost:  &cost,
		Price: &price,
	}
	mockErr := errors.New(mock.Anything)

	queryString :=
		"UPDATE `Topping` SET `name`=?,`description`=?,`cookingGuide`=?,`cost`=?,`price`=? WHERE id = ?"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *productmodel.ToppingUpdateInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name:   "Update topping successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: updateData,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(queryString).
					WithArgs(
						*updateData.Name,
						*updateData.Description,
						*updateData.CookingGuide,
						*updateData.Cost,
						*updateData.Price,
						id).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mockSqlDB.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:   "Update topping failed",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: updateData,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(queryString).
					WithArgs(
						*updateData.Name,
						*updateData.Description,
						*updateData.CookingGuide,
						*updateData.Cost,
						*updateData.Price,
						id).
					WillReturnError(mockErr)
				mockSqlDB.ExpectRollback()
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

			err := s.UpdateTopping(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateTopping() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateTopping() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
