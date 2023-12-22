package invoicestore

import (
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateInvoice(t *testing.T) {
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

	customerId := "customer001"
	invoiceData := &invoicemodel.InvoiceCreate{
		Id:                  "invoice123",
		CustomerId:          &customerId,
		TotalPrice:          100.0,
		AmountReceived:      100,
		AmountPriceUsePoint: 0.0,
		CreatedBy:           "user001",
	}

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *invoicemodel.InvoiceCreate
	}

	expectedSql := "INSERT INTO `Invoice` (`id`,`customerId`,`totalPrice`,`amountReceived`,`amountPriceUsePoint`,`createdBy`) VALUES (?,?,?,?,?,?)"
	expectedValues := []driver.Value{
		invoiceData.Id,
		invoiceData.CustomerId,
		invoiceData.TotalPrice,
		invoiceData.AmountReceived,
		invoiceData.AmountPriceUsePoint,
		invoiceData.CreatedBy,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create invoice in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: invoiceData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(expectedValues...).
					WillReturnResult(sqlmock.NewResult(0, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create invoice in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: invoiceData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(expectedValues...).
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

			err := s.CreateInvoice(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
