package cancelnotedetailstore

import (
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
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

func Test_sqlStore_CreateListCancelNoteDetail(t *testing.T) {
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
		data []cancelnotedetailmodel.CancelNoteDetailCreate
	}

	mockCancelNoteId := "123"
	mockReason := cancelnotedetailmodel.Damaged
	mockDetails := []cancelnotedetailmodel.CancelNoteDetailCreate{
		{
			CancelNoteId: mockCancelNoteId,
			IngredientId: "Ing001",
			ExpiryDate:   "12/12/2003",
			Reason:       &mockReason,
			AmountCancel: 100,
		},
		{
			CancelNoteId: mockCancelNoteId,
			IngredientId: "Ing002",
			ExpiryDate:   "11/12/2003",
			Reason:       &mockReason,
			AmountCancel: 100,
		},
	}
	mockErr := errors.New(mock.Anything)

	mockPlaceholders := make([]string, 0)
	mockArgsForQuery := make([]driver.Value, 0)
	for _, val := range mockDetails {
		mockPlaceholders = append(mockPlaceholders, "(?,?,?,?,?)")
		mockArgsForQuery = append(
			mockArgsForQuery,
			val.CancelNoteId,
			val.IngredientId,
			val.ExpiryDate,
			val.Reason,
			val.AmountCancel)
	}
	queryString := "INSERT INTO `CancelNoteDetail` " +
		"(`cancelNoteId`,`ingredientId`,`expiryDate`,`reason`,`amountCancel`)" +
		" VALUES " + strings.Join(mockPlaceholders, ",")

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create cancel note in database successfully",
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
			name: "Create cancel note in database failed",
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

			err := s.CreateListCancelNoteDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListCancelNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListCancelNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
