package importnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	filter2 "coffee_shop_management_backend/module/supplier/suppliermodel/filter"
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

func Test_sqlStore_ListImportNoteBySupplier(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
		DSN:                       "user:password@tcp(yourDBServer:port)/dbname?parseTime=true&loc=UTC",
	}), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(err)
	}

	supplierId := mock.Anything

	doneStatus := importnotemodel.Done
	cancelStatus := importnotemodel.Cancel
	unixNumber := int64(1702080000)
	timeByUnix := time.Date(2023, 12, 9, 0, 0, 0, 0, time.UTC)
	importNote1 := importnotemodel.ImportNote{
		Id:         "1",
		TotalPrice: 150,
		Status:     &doneStatus,
		CreatedAt:  &timeByUnix,
		SupplierId: "supplier1",
	}

	importNote2 := importnotemodel.ImportNote{
		Id:         "2",
		TotalPrice: 200.0,
		Status:     &cancelStatus,
		CreatedAt:  &timeByUnix,
		SupplierId: "supplier1",
	}

	filter := &filter2.SupplierImportFilter{
		DateFrom: &unixNumber,
		DateTo:   &unixNumber,
	}

	paging := &common.Paging{
		Page:  3,
		Limit: 5,
	}

	expectedSql := "SELECT * FROM `ImportNote` WHERE createdAt >= ? AND createdAt <= ? AND supplierId = ? LIMIT 5 OFFSET 10"
	expectedCountSql := "SELECT count(*) FROM `ImportNote` WHERE createdAt >= ? AND createdAt <= ? AND supplierId = ?"

	mockErr := errors.New(mock.Anything)
	var moreKeys []string
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx              context.Context
		supplierID       string
		filterImportNote *filter2.SupplierImportFilter
		paging           *common.Paging
		moreKeys         []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "List import notes by supplier successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:              context.Background(),
				supplierID:       supplierId,
				filterImportNote: filter,
				paging:           paging,
				moreKeys:         moreKeys,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), supplierId).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(12))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), supplierId).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "totalPrice", "status", "createdAt", "supplierId"}).
							AddRow(importNote1.Id, importNote1.TotalPrice, importNote1.Status, importNote1.CreatedAt, importNote1.SupplierId).
							AddRow(importNote2.Id, importNote2.TotalPrice, importNote2.Status, importNote2.CreatedAt, importNote2.SupplierId))
			},
			want:    []importnotemodel.ImportNote{importNote1, importNote2},
			wantErr: false,
		},
		{
			name: "List import notes by supplier failed because can not get total records of import notes",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:              context.Background(),
				supplierID:       supplierId,
				filterImportNote: filter,
				paging:           paging,
				moreKeys:         moreKeys,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), supplierId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List import notes by supplier failed because can not query import note records",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:              context.Background(),
				supplierID:       supplierId,
				filterImportNote: filter,
				paging:           paging,
				moreKeys:         moreKeys,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), supplierId).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), supplierId).
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

			got, err := s.ListImportNoteBySupplier(
				tt.args.ctx,
				tt.args.supplierID,
				tt.args.filterImportNote,
				tt.args.paging,
				tt.args.moreKeys...,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "ListImportNoteBySupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListImportNoteBySupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListImportNoteBySupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
