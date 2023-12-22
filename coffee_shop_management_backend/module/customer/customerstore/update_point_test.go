package customerstore

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_UpdateCustomerPoint(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	customerId := "1"
	amount := float32(50)
	updateData := &customermodel.CustomerUpdatePoint{
		Amount: &amount,
	}

	expectedSql := "UPDATE `Customer` SET `point`=point + ? WHERE id = ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *customermodel.CustomerUpdatePoint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update customer point successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(amount, customerId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update customer point failed because can not save to database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(amount, customerId).
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

			err := s.UpdateCustomerPoint(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateCustomerPoint() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateCustomerPoint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
