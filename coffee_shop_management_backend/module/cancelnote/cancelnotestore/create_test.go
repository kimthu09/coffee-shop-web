package cancelnotestore

import (
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateCancelNote(t *testing.T) {
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
		data *cancelnotemodel.CancelNoteCreate
	}

	cancelNoteId := "123"
	cancelNoteCreate := cancelnotemodel.CancelNoteCreate{
		Id:                      &cancelNoteId,
		TotalPrice:              0,
		CreateBy:                "123",
		CancelNoteCreateDetails: nil,
	}
	mockErr := errors.New(mock.Anything)
	expectedSql := "INSERT INTO `CancelNote` (`id`,`totalPrice`,`createBy`) VALUES (?,?,?)"

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
				data: &cancelNoteCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						*cancelNoteCreate.Id,
						cancelNoteCreate.TotalPrice,
						cancelNoteCreate.CreateBy).
					WillReturnResult(sqlmock.NewResult(1, 1))
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
				data: &cancelNoteCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						*cancelNoteCreate.Id,
						cancelNoteCreate.TotalPrice,
						cancelNoteCreate.CreateBy).
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

			err := s.CreateCancelNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateCancelNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateCancelNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
