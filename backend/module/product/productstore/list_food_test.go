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

func Test_sqlStore_ListFood(t *testing.T) {
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

	searchKey := "TestFood"
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
	listFood := []productmodel.Food{
		{
			Product: &productmodel.Product{
				Id:           "123",
				Name:         "TestFood",
				Description:  "Description for TestFood",
				CookingGuide: "CookingGuide for TestFood",
				IsActive:     true,
			},
		},
	}
	listFood2 := []productmodel.Food{
		{
			Product: &productmodel.Product{
				Id:           "123",
				Name:         "TestFood",
				Description:  "Description for TestFood",
				CookingGuide: "CookingGuide for TestFood",
				IsActive:     false,
			},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "description", "cookingGuide", "isActive"})
	for _, food := range listFood {
		rows.AddRow(food.Id, food.Name, food.Description, food.CookingGuide, food.IsActive)
	}
	rows2 := sqlmock.NewRows([]string{"id", "name", "description", "cookingGuide", "isActive"})
	for _, food := range listFood2 {
		rows2.AddRow(food.Id, food.Name, food.Description, food.CookingGuide, food.IsActive)
	}

	queryString :=
		"SELECT * FROM `Food` " +
			"WHERE (id LIKE ? OR name LIKE ?) " +
			"AND isActive = ? " +
			"ORDER BY name LIMIT 10"

	countRows := sqlmock.NewRows([]string{"count"})
	countRows.AddRow(len(listFood))
	queryStringCount :=
		"SELECT count(*) FROM `Food` " +
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
		want    []productmodel.Food
		wantErr bool
	}{
		{
			name:   "List food failed because can not get number of rows from database",
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
			want:    nil,
			wantErr: true,
		},
		{
			name:   "List food failed because can not get list from database",
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
			want:    nil,
			wantErr: true,
		},
		{
			name:   "List food successfully with active status",
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
			want:    listFood,
			wantErr: false,
		},
		{
			name:   "List food successfully with inactive status",
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
			want:    listFood2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListFood(
				tt.args.ctx,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListFood() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListFood() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
