package exportnotestore

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateExportNote(t *testing.T) {
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
		data *exportnotemodel.ExportNoteCreate
	}

	exportNoteId := "123"
	reason := exportnotemodel.OutOfDate
	exportNoteCreate := exportnotemodel.ExportNoteCreate{
		Id:        &exportNoteId,
		CreatedBy: "123",
		Reason:    &reason,
	}
	mockErr := errors.New(mock.Anything)
	expectedSql := "INSERT INTO `ExportNote` (`id`,`createdBy`,`reason`) VALUES (?,?,?)"

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
				data: &exportNoteCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						*exportNoteCreate.Id,
						exportNoteCreate.CreatedBy,
						*exportNoteCreate.Reason).
					WillReturnResult(sqlmock.NewResult(1, 1))
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
				data: &exportNoteCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						*exportNoteCreate.Id,
						exportNoteCreate.CreatedBy,
						*exportNoteCreate.Reason).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
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

			err := s.CreateExportNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateExportNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateExportNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
