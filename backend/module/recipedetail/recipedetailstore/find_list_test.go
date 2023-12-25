package recipedetailstore

import (
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_FindListRecipeDetail(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}

	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	recipeId := "Recipe001"
	mockData := []recipedetailmodel.RecipeDetail{
		{
			RecipeId:     recipeId,
			IngredientId: "Ingredient001",
			AmountNeed:   100,
		},
		{
			RecipeId:     recipeId,
			IngredientId: "Ingredient002",
			AmountNeed:   50,
		},
	}

	expectedQuery := "SELECT * FROM `RecipeDetail` WHERE `recipeId` = ?"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []recipedetailmodel.RecipeDetail
		wantErr bool
	}{
		{
			name: "Find recipe details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				conditions: map[string]interface{}{
					"recipeId": recipeId,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"recipeId", "ingredientId", "amountNeed"}).
					AddRow(recipeId, mockData[0].IngredientId, mockData[0].AmountNeed).
					AddRow(recipeId, mockData[1].IngredientId, mockData[1].AmountNeed)

				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(recipeId).
					WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Find recipe details failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				conditions: map[string]interface{}{
					"recipeId": recipeId,
				},
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(recipeId).
					WillReturnError(mockErr)
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

			got, err := s.FindListRecipeDetail(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindListRecipeDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindListRecipeDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindListRecipeDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
