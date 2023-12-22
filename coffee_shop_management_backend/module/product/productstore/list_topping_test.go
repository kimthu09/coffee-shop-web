package productstore

import (
	"coffee_shop_management_backend/common"
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

func Test_sqlStore_ListTopping(t *testing.T) {
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

	searchKey := "TestTopping"
	isActive := true
	inActive := false
	filter := &productmodel.Filter{
		SearchKey: searchKey,
		IsActive:  &isActive,
	}
	filter2 := &productmodel.Filter{
		SearchKey: searchKey,
		IsActive:  &inActive,
	}
	mockProperties := []string{"id", "name"}
	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}
	listTopping := []productmodel.Topping{
		{
			Product: &productmodel.Product{
				Id:           "123",
				Name:         "TestTopping",
				Description:  "Description for TestTopping",
				CookingGuide: "CookingGuide for TestTopping",
				IsActive:     true,
			},
			Cost:     10,
			Price:    15,
			RecipeId: "Recipe001",
		},
	}
	listTopping2 := []productmodel.Topping{
		{
			Product: &productmodel.Product{
				Id:           "123",
				Name:         "TestTopping",
				Description:  "Description for TestTopping",
				CookingGuide: "CookingGuide for TestTopping",
				IsActive:     false,
			},
			Cost:     10,
			Price:    15,
			RecipeId: "Recipe001",
		},
	}
	rows := sqlmock.NewRows([]string{
		"id", "name", "description", "cookingGuide", "isActive",
		"cost", "price", "recipeId"})
	for _, topping := range listTopping {
		rows.AddRow(topping.Id, topping.Name, topping.Description,
			topping.CookingGuide, topping.IsActive,
			topping.Cost, topping.Price, topping.RecipeId)
	}
	rows2 := sqlmock.NewRows([]string{
		"id", "name", "description", "cookingGuide", "isActive",
		"cost", "price", "recipeId"})
	for _, topping := range listTopping2 {
		rows2.AddRow(topping.Id, topping.Name, topping.Description,
			topping.CookingGuide, topping.IsActive,
			topping.Cost, topping.Price, topping.RecipeId)
	}

	queryString :=
		"SELECT * FROM `Topping` " +
			"WHERE (id LIKE ? OR name LIKE ?) " +
			"AND isActive = ? " +
			"ORDER BY name LIMIT 10"

	countRows := sqlmock.NewRows([]string{"count"})
	countRows.AddRow(len(listTopping))
	queryStringCount :=
		"SELECT count(*) FROM `Topping` " +
			"WHERE (id LIKE ? OR name LIKE ?) " +
			"AND isActive = ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *productmodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []productmodel.Topping
		wantErr bool
	}{
		{
			name:   "List topping failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						*filter.IsActive).
					WillReturnError(mockErr)
			},
			want:    listTopping,
			wantErr: true,
		},
		{
			name:   "List topping failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						*filter.IsActive).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						*filter.IsActive).
					WillReturnError(mockErr)
			},
			want:    listTopping,
			wantErr: true,
		},
		{
			name:   "List topping successfully with active status",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						*filter.IsActive).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						*filter.IsActive).
					WillReturnRows(rows)
			},
			want:    listTopping,
			wantErr: false,
		},
		{
			name:   "List topping successfully with inactive status",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter2,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter2.SearchKey+"%",
						"%"+filter2.SearchKey+"%",
						*filter2.IsActive).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter2.SearchKey+"%",
						"%"+filter2.SearchKey+"%",
						*filter2.IsActive).
					WillReturnRows(rows2)
			},
			want:    listTopping2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListTopping(
				tt.args.ctx,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListTopping() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListTopping() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListTopping() = %v, want %v", got, tt.want)
			}
		})
	}
}
