package importnotestore

import (
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateImportNote(t *testing.T) {
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
		data *importnotemodel.ImportNoteCreate
	}

	validId := "123456789012"
	importNoteCreate := importnotemodel.ImportNoteCreate{
		Id:         &validId,
		TotalPrice: 0,
		SupplierId: validId,
		CreateBy:   "123",
	}
	mockErr := errors.New(mock.Anything)
	expectedSql := "INSERT INTO `ImportNote` " +
		"(`id`,`totalPrice`,`supplierId`,`createBy`) VALUES (?,?,?,?)"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create import note in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						*importNoteCreate.Id,
						importNoteCreate.TotalPrice,
						importNoteCreate.SupplierId,
						importNoteCreate.CreateBy).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create import note in database failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &importNoteCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						*importNoteCreate.Id,
						importNoteCreate.TotalPrice,
						importNoteCreate.SupplierId,
						importNoteCreate.CreateBy).
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

			err := s.CreateImportNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateImportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateImportNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
