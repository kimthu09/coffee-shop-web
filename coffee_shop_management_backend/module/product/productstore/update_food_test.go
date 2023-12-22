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

func Test_sqlStore_UpdateFood(t *testing.T) {
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
	name := "UpdatedFoodName"
	description := "UpdatedDescription"
	cookingGuide := "UpdatedCookingGuide"
	updateData := &productmodel.FoodUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &name,
			Description:  &description,
			CookingGuide: &cookingGuide,
		},
	}
	mockErr := errors.New(mock.Anything)

	queryString :=
		"UPDATE `Food` SET `name`=?,`description`=?,`cookingGuide`=? WHERE id = ?"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *productmodel.FoodUpdateInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name:   "Update food successfully",
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
						id).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mockSqlDB.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:   "Update food failed",
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

			err := s.UpdateFood(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateFood() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateFood() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
