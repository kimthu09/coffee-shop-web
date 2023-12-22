package categorystore

import (
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

func Test_sqlStore_FindCategory(t *testing.T) {
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

	categoryId := mock.Anything
	categoryName := mock.Anything
	categoryDescription := mock.Anything
	category := categorymodel.Category{
		Id:          categoryId,
		Name:        categoryName,
		Description: categoryDescription,
	}
	conditions := map[string]interface{}{"id": categoryId}
	mockErr := errors.New("something went wrong with the database")

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	expectedSql := "SELECT * FROM `Category` WHERE `id` = ? ORDER BY `Category`.`id` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *categorymodel.Category
		wantErr bool
	}{
		{
			name: "Find category in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(categoryId).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "description"}).
							AddRow(category.Id, category.Name, category.Description))
			},
			want:    &category,
			wantErr: false,
		},
		{
			name: "Find category in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(categoryId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find category in database record not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(categoryId).
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

			got, err := s.FindCategory(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindCategory() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindCategory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
