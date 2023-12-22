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

func Test_sqlStore_CreateFood(t *testing.T) {
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

	id := "product001"
	foodCreate := &productmodel.FoodCreate{
		ProductCreate: &productmodel.ProductCreate{
			Id:           &id,
			Name:         "TestFood",
			Description:  "Description for TestFood",
			CookingGuide: "Cooking guide for TestFood",
		},
	}
	expectedQuery := "INSERT INTO `Food` (`id`,`name`,`description`,`cookingGuide`) VALUES (?,?,?,?)"
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *productmodel.FoodCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name:   "Create food successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: foodCreate,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WithArgs(
						foodCreate.Id,
						foodCreate.Name,
						foodCreate.Description,
						foodCreate.CookingGuide,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockSqlDB.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:   "Create food failed because can not save data to database",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: foodCreate,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WithArgs(
						foodCreate.Id,
						foodCreate.Name,
						foodCreate.Description,
						foodCreate.CookingGuide,
					).
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

			err := s.CreateFood(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
