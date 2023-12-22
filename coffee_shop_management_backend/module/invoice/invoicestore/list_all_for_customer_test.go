package invoicestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_sqlStore_ListAllInvoiceByCustomer(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	customerId := "customer123"
	dateFrom := int64(1709500430)
	dateTo := int64(1709500432)
	dateFromTime := time.Unix(dateFrom, 0)
	dateToTime := time.Unix(dateTo, 0)

	filter := customermodel.FilterInvoice{
		DateFrom: &dateFrom,
		DateTo:   &dateTo,
	}

	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}

	listInvoice := []invoicemodel.Invoice{
		{
			Id:         "invoice123",
			CustomerId: customerId,
			TotalPrice: 100.0,
			CreatedAt:  &time.Time{},
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"customerId",
		"totalPrice",
		"createdAt",
	})
	for _, invoice := range listInvoice {
		rows.AddRow(
			invoice.Id,
			invoice.CustomerId,
			invoice.TotalPrice,
			invoice.CreatedAt,
		)
	}

	queryString :=
		"SELECT * FROM `Invoice` " +
			"WHERE customerId = ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"ORDER BY createdAt desc LIMIT 10"

	countRows := sqlmock.NewRows([]string{"count"})
	countRows.AddRow(len(listInvoice))
	queryStringCount :=
		"SELECT count(*) FROM `Invoice` " +
			"WHERE customerId = ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		customerID string
		filter     *customermodel.FilterInvoice
		paging     *common.Paging
		moreKeys   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []invoicemodel.Invoice
		wantErr bool
	}{
		{
			name:   "List invoices by customer failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				customerID: customerId,
				filter:     &filter,
				paging:     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						customerId,
						dateFromTime,
						dateToTime,
					).
					WillReturnError(mockErr)
			},
			want:    listInvoice,
			wantErr: true,
		},
		{
			name:   "List invoices by customer failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				customerID: customerId,
				filter:     &filter,
				paging:     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						customerId,
						dateFromTime,
						dateToTime,
					).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						customerId,
						dateFromTime,
						dateToTime,
					).
					WillReturnError(mockErr)
			},
			want:    listInvoice,
			wantErr: true,
		},
		{
			name:   "List invoices by customer successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				customerID: customerId,
				filter:     &filter,
				paging:     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						customerId,
						dateFromTime,
						dateToTime,
					).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						customerId,
						dateFromTime,
						dateToTime,
					).
					WillReturnRows(rows)
			},
			want:    listInvoice,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListAllInvoiceByCustomer(
				tt.args.ctx,
				tt.args.customerID,
				tt.args.filter,
				tt.args.paging,
				tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "ListAllInvoiceByCustomer() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListAllInvoiceByCustomer() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListAllInvoiceByCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
