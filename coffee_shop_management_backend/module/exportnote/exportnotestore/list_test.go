package exportnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
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

func Test_sqlStore_ListExportNote(t *testing.T) {
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

	minPrice := float32(10)
	maxPrice := float32(90)
	filter := exportnotemodel.Filter{
		SearchKey: "mockSearchKey",
		MinPrice:  &minPrice,
		MaxPrice:  &maxPrice,
	}
	mockProperties := []string{"id", "createBy"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	listExportNote := []exportnotemodel.ExportNote{
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
	for _, exportNote := range listExportNote {
		rows.AddRow(exportNote.Id, exportNote.TotalPrice, exportNote.CreateAt, exportNote.CreateBy)
	}
	queryString := "SELECT * FROM `ExportNote` WHERE " +
		"(id LIKE ? OR createBy LIKE ?) AND" +
		" totalPrice >= ? AND totalPrice <= ? " +
		"ORDER BY createAt desc LIMIT " + strconv.FormatInt(paging.Limit, 10)

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listExportNote))
	queryStringCount := "SELECT count(*) FROM `ExportNote` WHERE " +
		"(id LIKE ? OR createBy LIKE ?) AND " +
		"totalPrice >= ? AND totalPrice <= ? "

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *exportnotemodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []exportnotemodel.ExportNote
		wantErr bool
	}{
		{
			name:   "List export note failed because can not get number of rows from database",
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
			want:    listExportNote,
			wantErr: true,
		},
		{
			name:   "List export note failed because can not get list from database",
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
			want:    listExportNote,
			wantErr: true,
		},
		{
			name:   "List export note successfully",
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
			want:    listExportNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListExportNote(
				tt.args.ctx,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListExportNote() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListExportNote() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListExportNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
