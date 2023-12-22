package categoryfoodstore

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

func Test_sqlStore_FindListCategories(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}

	foodId := "1"
	category1 := categorymodel.SimpleCategoryWithId{
		CategoryId: "1",
	}
	category2 := categorymodel.SimpleCategoryWithId{
		CategoryId: "2",
	}

	expectedSql := "SELECT * FROM `CategoryFood` WHERE foodId = ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx    context.Context
		foodId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []categorymodel.SimpleCategoryWithId
		wantErr bool
	}{
		{
			name: "Find list of categories for food successfully",
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
					WillReturnRows(sqlmock.NewRows([]string{"categoryId"}).
						AddRow(category1.CategoryId).
						AddRow(category2.CategoryId))
			},
			want:    []categorymodel.SimpleCategoryWithId{category1, category2},
			wantErr: false,
		},
		{
			name: "Find list of categories for food not found",
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
					WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find list of categories for food failed something went wrong with the database",
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

			got, err := s.FindListCategories(tt.args.ctx, tt.args.foodId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindListCategories() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindListCategories() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindListCategories() got = %v, want %v", got, tt.want)
			}
		})
	}
}
