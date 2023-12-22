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

func Test_sqlStore_GetAllIngredient(t *testing.T) {
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

	measureType := enum.Volume
	mockData := []ingredientmodel.SimpleIngredient{
		{
			Id:          "Ingredient001",
			Name:        "IngredientName001",
			MeasureType: &measureType,
		},
		{
			Id:          "Ingredient002",
			Name:        "IngredientName002",
			MeasureType: &measureType,
		},
	}

	expectedQuery := "SELECT * FROM `Ingredient`"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		mock    func()
		want    []ingredientmodel.SimpleIngredient
		wantErr bool
	}{
		{
			name: "Get all ingredients successfully",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "measureType"})

				for _, data := range mockData {
					rows.AddRow(data.Id, data.Name, data.MeasureType)
				}

				sqlDBMock.ExpectQuery(expectedQuery).WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Get all ingredients failed",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).WillReturnError(mockErr)
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

			got, err := s.GetAllIngredient(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err, "GetAllIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetAllIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got, "GetAllIngredient() got = %v, want %v", got, tt.want)
		})
	}
}
