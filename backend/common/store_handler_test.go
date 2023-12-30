package common

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGetWhereClause(t *testing.T) {
	sqlDB, _, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	type args struct {
		db                         *gorm.DB
		searchKey                  string
		propertiesContainSearchKey []string
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
		{
			name: "Get DB has been searched by properties successfully",
			args: args{
				db:                         gormDB,
				searchKey:                  "search",
				propertiesContainSearchKey: []string{"name", "id"},
			},
			want: gormDB.Where("name LIKE ? OR id LIKE ?", "%search%", "%search%"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetWhereClause(tt.args.db, tt.args.searchKey, tt.args.propertiesContainSearchKey)

			assert.Equal(
				t,
				tt.want,
				got,
				"GetWhereClause() = %v, want %v", got, tt.want)
		})
	}
}

func TestHandlePaging(t *testing.T) {
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

	type args struct {
		db     *gorm.DB
		paging *Paging
	}

	gormDB = gormDB.Table("Users")
	copiedGormDB1 := gormDB.Session(&gorm.Session{})
	copiedGormDB2 := gormDB.Session(&gorm.Session{})
	expectedCountSql := "SELECT count(*) FROM `Users`"

	tests := []struct {
		name      string
		args      args
		mock      func()
		want      *gorm.DB
		wantTotal int64
		wantErr   bool
	}{
		{
			name: "Get DB has been handled paging failed because can not get the total",
			args: args{
				db: copiedGormDB1,
				paging: &Paging{
					Page:  2,
					Limit: 10,
				},
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedCountSql).
					WithArgs().
					WillReturnError(errors.New(mock.Anything))
			},
			wantTotal: 0,
			want:      nil,
			wantErr:   true,
		},
		{
			name: "Get DB has been handled paging successfully",
			args: args{
				db: copiedGormDB2,
				paging: &Paging{
					Page:  2,
					Limit: 10,
				},
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedCountSql).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(11))
			},
			wantTotal: 11,
			want:      copiedGormDB2.Offset(10).Limit(10),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := HandlePaging(tt.args.db, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"HandlePaging() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"HandlePaging() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"HandlePaging() = %v, want %v", got, tt.want)
				assert.Equal(
					t,
					tt.wantTotal,
					tt.args.paging.Total,
					"HandlePaging() total = %v, want %v", tt.args.paging.Total, tt.wantTotal)
			}
		})
	}
}
