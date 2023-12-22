package inventorychecknotestore

import (
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateInventoryCheckNote(t *testing.T) {
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
		data *inventorychecknotemodel.InventoryCheckNoteCreate
	}

	idNote := "123"
	mockData := &inventorychecknotemodel.InventoryCheckNoteCreate{
		Id:                &idNote,
		AmountDifferent:   10,
		AmountAfterAdjust: 20,
		CreatedBy:         "user123",
	}
	mockErr := errors.New(mock.Anything)
	expectedQuery := "INSERT INTO `InventoryCheckNote` (`id`,`amountDifferent`,`amountAfterAdjust`,`createdBy`) VALUES (?,?,?,?)"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create inventory check note successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedQuery).
					WithArgs(
						mockData.Id,
						mockData.AmountDifferent,
						mockData.AmountAfterAdjust,
						mockData.CreatedBy).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create inventory check note failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedQuery).
					WithArgs(
						mockData.Id,
						mockData.AmountDifferent,
						mockData.AmountAfterAdjust,
						mockData.CreatedBy).
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

			err := s.CreateInventoryCheckNote(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
