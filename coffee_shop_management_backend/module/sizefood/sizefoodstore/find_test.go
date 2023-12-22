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

func Test_sqlStore_FindSizeFood(t *testing.T) {
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

	foodId := "Food001"
	sizeId := "Size001"
	recipeId := "Recipe001"
	mockConditions := map[string]interface{}{
		"foodId": foodId,
		"sizeId": sizeId,
	}
	sizeFood := sizefoodmodel.SizeFood{
		FoodId:   foodId,
		SizeId:   sizeId,
		Name:     mock.Anything,
		Cost:     0,
		Price:    0,
		RecipeId: recipeId,
	}
	expectedQuery := "SELECT * FROM `SizeFood` WHERE `foodId` = ? AND `sizeId` = ? ORDER BY `SizeFood`.`foodId` LIMIT 1"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *sizefoodmodel.SizeFood
		wantErr bool
	}{
		{
			name: "Find size food in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedQuery).
					WithArgs(mockConditions["foodId"], mockConditions["sizeId"]).
					WillReturnRows(
						sqlmock.NewRows([]string{"foodId", "sizeId", "name", "price", "cost", "recipeId"}).
							AddRow(mockConditions["foodId"], mockConditions["sizeId"], sizeFood.Name,
								sizeFood.Price, sizeFood.Cost, sizeFood.RecipeId))
			},
			want:    &sizeFood,
			wantErr: false,
		},
		{
			name: "Find size food in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedQuery).
					WithArgs(mockConditions["foodId"], mockConditions["sizeId"]).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find size food in database record not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedQuery).
					WithArgs(mockConditions["foodId"], mockConditions["sizeId"]).
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

			got, err := s.FindSizeFood(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindSizeFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindSizeFood() got = %v, want %v", got, tt.want)
			}
		})
	}
}
