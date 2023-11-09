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

func Test_sqlStore_UpdateImportNote(t *testing.T) {
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

	validId := "123"
	expectedSql := "UPDATE `ImportNote` " +
		"SET `closeBy`=?,`status`=? " +
		"WHERE id = ?"
	status := importnotemodel.Done
	dataUpdate := importnotemodel.ImportNoteUpdate{
		CloseBy:    mock.Anything,
		Id:         validId,
		SupplierId: mock.Anything,
		TotalPrice: 0,
		Status:     &status,
	}
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *importnotemodel.ImportNoteUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update import note successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   validId,
				data: &dataUpdate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						dataUpdate.CloseBy,
						dataUpdate.Status,
						validId,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update import note failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   validId,
				data: &dataUpdate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						dataUpdate.CloseBy,
						dataUpdate.Status,
						validId,
					).
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

			err := s.UpdateImportNote(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateImportNote() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateImportNote() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
