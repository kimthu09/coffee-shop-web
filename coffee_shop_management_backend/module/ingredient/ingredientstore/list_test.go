package ingredientstore

import (
	"coffee_shop_management_backend/common"
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

func Test_sqlStore_ListIngredient(t *testing.T) {
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
		ctx                        context.Context
		filter                     *ingredientmodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}

	searchKey := "Ingredient00"
	minPrice := float32(0)
	maxPrice := float32(1000)
	minAmount := 0
	maxAmount := 1000
	measureType := enum.Unit
	filter := ingredientmodel.Filter{
		SearchKey:   searchKey,
		MinPrice:    &minPrice,
		MaxPrice:    &maxPrice,
		MinAmount:   &minAmount,
		MaxAmount:   &maxAmount,
		MeasureType: "Unit",
	}
	mockProperties := []string{"id", "name"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}

	mockData := []ingredientmodel.Ingredient{
		{
			Id:          "Ingredient001",
			Name:        "IngredientName001",
			Amount:      10,
			MeasureType: &measureType,
			Price:       5.0,
		},
		{
			Id:          "Ingredient002",
			Name:        "IngredientName002",
			Amount:      20,
			MeasureType: &measureType,
			Price:       3.0,
		},
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"amount",
		"measureType",
		"price",
	})
	for _, ingredient := range mockData {
		rows.AddRow(
			ingredient.Id,
			ingredient.Name,
			ingredient.Amount,
			ingredient.MeasureType,
			ingredient.Price)
	}

	queryString :=
		"SELECT id, name, amount, measureType, price FROM `Ingredient` " +
			"WHERE (id LIKE ? AND name LIKE ?) " +
			"AND price >= ? " +
			"AND price <= ? " +
			"AND amount >= ? " +
			"AND amount <= ? " +
			"AND measureType = ? " +
			"ORDER BY name LIMIT 10"

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(mockData))
	queryStringCount :=
		"SELECT count(*) FROM `Ingredient` " +
			"WHERE (id LIKE ? OR name LIKE ?) " +
			"AND measureType = ? " +
			"AND price >= ? " +
			"AND price <= ? " +
			"AND amount >= ? " +
			"AND amount <= ?"

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		mock    func()
		args    args
		want    []ingredientmodel.Ingredient
		wantErr bool
	}{
		{
			name:   "List ingredient failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						measureType,
						minPrice,
						maxPrice,
						minAmount,
						maxAmount,
					).
					WillReturnError(mockErr)
			},
			want:    mockData,
			wantErr: true,
		},
		{
			name:   "List ingredient failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						measureType,
						minPrice,
						maxPrice,
						minAmount,
						maxAmount,
					).
					WillReturnError(nil)

				sqlDBMock.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						measureType,
						minPrice,
						maxPrice,
						minAmount,
						maxAmount,
					).
					WillReturnError(mockErr)
			},
			want:    mockData,
			wantErr: true,
		},
		{
			name:   "List successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						measureType,
						minPrice,
						maxPrice,
						minAmount,
						maxAmount,
					).
					WillReturnError(nil)

				sqlDBMock.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						measureType,
						minPrice,
						maxPrice,
						minAmount,
						maxAmount,
					).
					WillReturnError(nil)
			},
			want:    mockData,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListIngredient(tt.args.ctx, tt.args.filter, tt.args.propertiesContainSearchKey, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListIngredient() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListIngredient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
