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

func Test_sqlStore_FindFood(t *testing.T) {
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

	foodId := "123"
	foodName := "TestFood"
	foodDescription := "Description for TestFood"
	foodCookingGuide := "Coking Guide for TestFood"
	foodActive := true
	conditions := map[string]interface{}{"id": foodId}
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	expectedSql := "SELECT * FROM `Food` WHERE `id` = ? ORDER BY `Food`.`id` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Food
		wantErr bool
	}{
		{
			name:   "Find food in database successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedSql).
					WithArgs(foodId).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "description", "cookingGuide", "isActive"}).
							AddRow(foodId, foodName, foodDescription, foodCookingGuide, foodActive))
			},
			want: &productmodel.Food{
				Product: &productmodel.Product{
					Id:           foodId,
					Name:         foodName,
					Description:  foodDescription,
					CookingGuide: foodCookingGuide,
					IsActive:     true,
				},
			},
			wantErr: false,
		},
		{
			name:   "Find food in database failed due to something went wrong with the database",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedSql).
					WithArgs(foodId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Find food in database record not found",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedSql).
					WithArgs(foodId).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindFood(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindFood() got = %v, want %v", got, tt.want)
			}
		})
	}
}
