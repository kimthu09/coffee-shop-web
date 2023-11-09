package exportnotedetailstore

import (
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"context"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func Test_sqlStore_CreateListExportNoteDetail(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data []exportnotedetailmodel.ExportNoteDetailCreate
	}

	mockExportNoteId := "123"
	mockDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			ExportNoteId: mockExportNoteId,
			IngredientId: "Ing001",
			ExpiryDate:   "12/12/2003",
			AmountExport: 100,
		},
		{
			ExportNoteId: mockExportNoteId,
			IngredientId: "Ing002",
			ExpiryDate:   "11/12/2003",
			AmountExport: 100,
		},
	}
	mockErr := errors.New(mock.Anything)

	mockPlaceholders := make([]string, 0)
	mockArgsForQuery := make([]driver.Value, 0)
	for _, val := range mockDetails {
		mockPlaceholders = append(mockPlaceholders, "(?,?,?,?)")
		mockArgsForQuery = append(
			mockArgsForQuery,
			val.ExportNoteId,
			val.IngredientId,
			val.ExpiryDate,
			val.AmountExport)
	}
	queryString := "INSERT INTO `ExportNoteDetail` " +
		"(`exportNoteId`,`ingredientId`,`expiryDate`,`amountExport`)" +
		" VALUES " + strings.Join(mockPlaceholders, ",")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create export note in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockDetails,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(queryString).
					WithArgs(mockArgsForQuery...).
					WillReturnResult(sqlmock.NewResult(1, int64(len(mockDetails))))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create export note in database failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockDetails,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(queryString).
					WithArgs(mockArgsForQuery...).
					WillReturnError(mockErr)
				sqlDBMock.ExpectCommit()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateListExportNoteDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListExportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListExportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
