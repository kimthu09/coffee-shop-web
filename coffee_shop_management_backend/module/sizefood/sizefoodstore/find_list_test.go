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

func Test_sqlStore_FindListSizeFood(t *testing.T) {
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
		ctx    context.Context
		foodId string
	}

	foodId := mock.Anything
	expectedSql := "SELECT * FROM `SizeFood` WHERE foodId = ?"
	sizeFoods := []sizefoodmodel.SizeFood{
		{
			FoodId:   foodId,
			SizeId:   mock.Anything,
			Name:     mock.Anything,
			Cost:     0,
			Price:    0,
			RecipeId: mock.Anything,
		},
		{
			FoodId:   foodId,
			SizeId:   mock.Anything,
			Name:     mock.Anything,
			Cost:     0,
			Price:    0,
			RecipeId: mock.Anything,
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []sizefoodmodel.SizeFood
		wantErr bool
	}{
		{
			name: "Find list size food by food id successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:    context.Background(),
				foodId: foodId,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(foodId).
					WillReturnRows(
						sqlmock.
							NewRows([]string{"foodId", "sizeId", "name", "price", "cost", "recipeId"}).
							AddRow(sizeFoods[0].FoodId, sizeFoods[0].SizeId,
								sizeFoods[0].Name, sizeFoods[0].Price,
								sizeFoods[0].Cost, sizeFoods[0].RecipeId).
							AddRow(sizeFoods[1].FoodId, sizeFoods[1].SizeId,
								sizeFoods[1].Name, sizeFoods[1].Price,
								sizeFoods[1].Cost, sizeFoods[1].RecipeId),
					)
			},
			want:    sizeFoods,
			wantErr: false,
		},
		{
			name: "Find list feature by role failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:    context.Background(),
				foodId: foodId,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(foodId).
					WillReturnError(mockErr)
			},
			want:    sizeFoods,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindListSizeFood(tt.args.ctx, tt.args.foodId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindListSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindListSizeFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindListSizeFood() got = %v, want %v", got, tt.want)
			}
		})
	}
}
