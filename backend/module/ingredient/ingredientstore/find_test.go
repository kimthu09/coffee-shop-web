package ingredientstore

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_FindIngredient(t *testing.T) {
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

	measureType := enum.Weight
	ingredientId := "Ingredient001"
	mockData := &ingredientmodel.Ingredient{
		Id:          ingredientId,
		Name:        "IngredientName001",
		Amount:      10,
		MeasureType: &measureType,
		Price:       5.0,
	}

	expectedQuery := "SELECT * FROM `Ingredient` WHERE `id` = ? ORDER BY `Ingredient`.`id` LIMIT 1"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *ingredientmodel.Ingredient
		wantErr bool
	}{
		{
			name: "Find ingredient successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				conditions: map[string]interface{}{
					"id": ingredientId,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "amount", "measureType", "price"}).
					AddRow(mockData.Id, mockData.Name, mockData.Amount, mockData.MeasureType, mockData.Price)

				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(ingredientId).
					WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Find ingredient failed because can not get data",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				conditions: map[string]interface{}{
					"id": ingredientId,
				},
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(ingredientId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find ingredient failed because can not find ingredient in the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: map[string]interface{}{"id": ingredientId},
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(ingredientId).
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

			got, err := s.FindIngredient(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindIngredient() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindIngredient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
