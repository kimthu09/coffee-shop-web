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

func Test_sqlStore_CreateListRecipeDetail(t *testing.T) {
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
		ctx  context.Context
		data []recipedetailmodel.RecipeDetailCreate
	}

	mockData := []recipedetailmodel.RecipeDetailCreate{
		{
			RecipeId:     "Recipe001",
			IngredientId: "Ingredient001",
			AmountNeed:   100,
		},
		{
			RecipeId:     "Recipe001",
			IngredientId: "Ingredient002",
			AmountNeed:   50,
		},
	}

	expectedQuery := "INSERT INTO `RecipeDetail` (`recipeId`,`ingredientId`,`amountNeed`) VALUES (?,?,?),(?,?,?)"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create recipe details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(
						mockData[0].RecipeId, mockData[0].IngredientId, mockData[0].AmountNeed,
						mockData[1].RecipeId, mockData[1].IngredientId, mockData[1].AmountNeed,
					).
					WillReturnResult(sqlmock.NewResult(1, int64(len(mockData))))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create recipe details failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(
						mockData[0].RecipeId, mockData[0].IngredientId, mockData[0].AmountNeed,
						mockData[1].RecipeId, mockData[1].IngredientId, mockData[1].AmountNeed,
					).
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

			err := s.CreateListRecipeDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListRecipeDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListRecipeDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
