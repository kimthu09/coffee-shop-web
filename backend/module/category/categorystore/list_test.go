package categorystore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_ListCategory(t *testing.T) {
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

	category1 := categorymodel.Category{
		Id:          "1",
		Name:        "Category1",
		Description: "Description1",
	}
	category2 := categorymodel.Category{
		Id:          "2",
		Name:        "Category2",
		Description: "Description2",
	}

	filter := &categorymodel.Filter{
		SearchKey: "Category",
	}

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	expectedSql := "SELECT * FROM `Category` WHERE name LIKE ? OR description LIKE ? ORDER BY name LIMIT 10"
	expectedCountSql := "SELECT count(*) FROM `Category` WHERE name LIKE ? OR description LIKE ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *categorymodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []categorymodel.Category
		wantErr bool
	}{
		{
			name: "List categories successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter,
				propertiesContainSearchKey: []string{"name", "description"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs("%Category%", "%Category%").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs("%Category%", "%Category%").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
						AddRow(category1.Id, category1.Name, category1.Description).
						AddRow(category2.Id, category2.Name, category2.Description))
			},
			want:    []categorymodel.Category{category1, category2},
			wantErr: false,
		},
		{
			name: "List categories failed because can get total records of categories",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter,
				propertiesContainSearchKey: []string{"name", "description"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs("%Category%", "%Category%").
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List categories failed because can not query category records",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filter,
				propertiesContainSearchKey: []string{"name", "description"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs("%Category%", "%Category%").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs("%Category%", "%Category%").
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

			got, err := s.ListCategory(tt.args.ctx, tt.args.filter, tt.args.propertiesContainSearchKey, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListCategory() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
