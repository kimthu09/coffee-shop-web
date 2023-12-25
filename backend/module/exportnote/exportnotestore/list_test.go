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
	"testing"
	"time"
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

	dateInt := int64(123)
	dateTime := time.Unix(dateInt, 0)
	createdBy := "user001"
	reason := exportnotemodel.OutOfDate
	filter := exportnotemodel.Filter{
		SearchKey:         "mockSearchKey",
		DateFromCreatedAt: &dateInt,
		DateToCreatedAt:   &dateInt,
		CreatedBy:         &createdBy,
		Reason:            &reason,
	}
	mockProperties := []string{"id"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	listExportNote := []exportnotemodel.ExportNote{
		{
			Id:        "123",
			Reason:    &reason,
			CreatedAt: &dateTime,
			CreatedBy: createdBy,
		},
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"reason",
		"createdAt",
		"createdBy",
	})
	for _, exportNote := range listExportNote {
		rows.AddRow(exportNote.Id, exportNote.Reason, exportNote.CreatedAt, exportNote.CreatedBy)
	}

	queryString :=
		"SELECT `ExportNote`.`id`,`ExportNote`.`reason`,`ExportNote`.`createdAt`,`ExportNote`.`createdBy` FROM `ExportNote` " +
			"JOIN MUser ON ExportNote.createdBy = MUser.id " +
			"WHERE id LIKE ? " +
			"AND reason = ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"AND MUser.Id = ? " +
			"ORDER BY createdAt desc LIMIT 10"

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listExportNote))
	queryStringCount :=
		"SELECT count(*) FROM `ExportNote` " +
			"JOIN MUser ON ExportNote.createdBy = MUser.id " +
			"WHERE id LIKE ? " +
			"AND reason = ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"AND MUser.Id = ?"

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
						reason,
						dateTime,
						dateTime,
						createdBy).
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
						reason,
						dateTime,
						dateTime,
						createdBy).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						reason,
						dateTime,
						dateTime,
						createdBy).
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
						reason,
						dateTime,
						dateTime,
						createdBy).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						reason,
						dateTime,
						dateTime,
						createdBy).
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
