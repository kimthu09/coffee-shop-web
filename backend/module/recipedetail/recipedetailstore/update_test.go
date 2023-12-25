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

func Test_sqlStore_UpdateRecipeDetail(t *testing.T) {
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
		ctx          context.Context
		idRecipe     string
		idIngredient string
		data         *recipedetailmodel.RecipeDetailUpdate
	}

	expectedQuery := "UPDATE `RecipeDetail` SET `amountNeed`=? WHERE recipeId = ? and ingredientId = ?"
	mockErr := errors.New(mock.Anything)

	recipeId := "Recipe001"
	ingredientId := "Ingredient001"

	mockData := &recipedetailmodel.RecipeDetailUpdate{
		IngredientId: ingredientId,
		AmountNeed:   75.0,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update recipe detail successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:          context.Background(),
				idRecipe:     recipeId,
				idIngredient: ingredientId,
				data:         mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.AmountNeed, recipeId, ingredientId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update recipe detail failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:          context.Background(),
				idRecipe:     recipeId,
				idIngredient: ingredientId,
				data:         mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.AmountNeed, recipeId, ingredientId).
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

			err := s.UpdateRecipeDetail(tt.args.ctx, tt.args.idRecipe, tt.args.idIngredient, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRecipeDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateRecipeDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
