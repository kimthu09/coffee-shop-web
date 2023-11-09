package cancelnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func Test_sqlStore_ListCancelNote(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *cancelnotemodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}

	minPrice := float32(10)
	maxPrice := float32(90)
	filter := cancelnotemodel.Filter{
		SearchKey: "mockSearchKey",
		MinPrice:  &minPrice,
		MaxPrice:  &maxPrice,
	}
	mockProperties := []string{"id", "createBy"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	listCancelNote := []cancelnotemodel.CancelNote{
		{
			Id:         "123",
			TotalPrice: 12000,
			CreateAt:   nil,
			CreateBy:   "123",
		},
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"totalPrice",
		"createAt",
		"createBy",
	})
	for _, cancelNote := range listCancelNote {
		rows.AddRow(cancelNote.Id, cancelNote.TotalPrice, cancelNote.CreateAt, cancelNote.CreateBy)
	}
	queryString := "SELECT * FROM `CancelNote` WHERE " +
		"(id LIKE ? OR createBy LIKE ?) AND" +
		" totalPrice >= ? AND totalPrice <= ? " +
		"ORDER BY createAt desc LIMIT " + strconv.FormatInt(paging.Limit, 10)

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listCancelNote))
	queryStringCount := "SELECT count(*) FROM `CancelNote` WHERE " +
		"(id LIKE ? OR createBy LIKE ?) AND " +
		"totalPrice >= ? AND totalPrice <= ? "

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []cancelnotemodel.CancelNote
		wantErr bool
	}{
		{
			name:   "List cancel note failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.MinPrice,
						filter.MaxPrice).
					WillReturnError(mockErr)
			},
			want:    listCancelNote,
			wantErr: true,
		},
		{
			name:   "List cancel note failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.MinPrice,
						filter.MaxPrice).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.MinPrice,
						filter.MaxPrice).
					WillReturnError(mockErr)
			},
			want:    listCancelNote,
			wantErr: true,
		},
		{
			name:   "List cancel note successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.MinPrice,
						filter.MaxPrice).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.MinPrice,
						filter.MaxPrice).
					WillReturnRows(rows)
			},
			want:    listCancelNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListCancelNote(tt.args.ctx, tt.args.filter, tt.args.propertiesContainSearchKey, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListCancelNote() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListCancelNote() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListCancelNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
