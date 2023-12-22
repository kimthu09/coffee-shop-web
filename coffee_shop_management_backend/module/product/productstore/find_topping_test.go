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

func Test_sqlStore_FindTopping(t *testing.T) {
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

	toppingId := "123"
	toppingName := "TestTopping"
	toppingDescription := "Description for TestTopping"
	toppingCookingGuide := "Coking Guide for TestTopping"
	toppingActive := true
	toppingCost := 10
	toppingPrice := 8
	toppingRecipeId := "Recipe001"
	conditions := map[string]interface{}{"id": toppingId}
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	expectedSql := "SELECT * FROM `Topping` WHERE `id` = ? ORDER BY `Topping`.`id` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Topping
		wantErr bool
	}{
		{
			name:   "Find topping in database successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedSql).
					WithArgs(toppingId).
					WillReturnRows(
						sqlmock.NewRows([]string{
							"id", "name", "description", "cookingGuide", "isActive", "cost", "price", "recipeId",
						}).AddRow(toppingId, toppingName, toppingDescription,
							toppingCookingGuide, toppingActive,
							toppingCost, toppingPrice, toppingRecipeId))
			},
			want: &productmodel.Topping{
				Product: &productmodel.Product{
					Id:           toppingId,
					Name:         toppingName,
					Description:  toppingDescription,
					CookingGuide: toppingCookingGuide,
					IsActive:     true,
				},
				Cost:     toppingCost,
				Price:    toppingPrice,
				RecipeId: toppingRecipeId,
			},
			wantErr: false,
		},
		{
			name:   "Find topping in database failed due to something went wrong with the database",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedSql).
					WithArgs(toppingId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Find topping in database record not found",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedSql).
					WithArgs(toppingId).
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

			got, err := s.FindTopping(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindTopping() got = %v, want %v", got, tt.want)
			}
		})
	}
}
