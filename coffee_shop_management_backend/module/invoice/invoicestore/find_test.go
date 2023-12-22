package invoicestore

import (
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_sqlStore_FindInvoice(t *testing.T) {
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

	invoiceData := &invoicemodel.Invoice{
		Id:                  "invoice123",
		CustomerId:          "customer456",
		TotalPrice:          100.0,
		AmountReceived:      100,
		AmountPriceUsePoint: 0.0,
		CreatedBy:           "user001",
		CreatedAt:           &time.Time{},
	}

	conditions := map[string]interface{}{"id": "invoice123"}
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	expectedSql := "SELECT * FROM `Invoice` WHERE `id` = ? ORDER BY `Invoice`.`id` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *invoicemodel.Invoice
		wantErr bool
	}{
		{
			name: "Find invoice in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(conditions["id"]).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "customerId", "totalPrice", "amountReceived", "amountPriceUsePoint", "createdBy", "createdAt"}).
							AddRow(invoiceData.Id, invoiceData.CustomerId, invoiceData.TotalPrice, invoiceData.AmountReceived, invoiceData.AmountPriceUsePoint, invoiceData.CreatedBy, invoiceData.CreatedAt))
			},
			want:    invoiceData,
			wantErr: false,
		},
		{
			name: "Find invoice in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(conditions["id"]).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find invoice in database record not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(conditions["id"]).
					WillReturnError(gorm.ErrRecordNotFound)
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

			got, err := s.FindInvoice(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindInvoice() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindInvoice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
