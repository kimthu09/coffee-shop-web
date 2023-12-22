package invoicestore

import (
	"coffee_shop_management_backend/common"
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

func Test_sqlStore_ListInvoice(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	dateInt := int64(123)
	dateTime := time.Unix(dateInt, 0)
	createdBy := "user001"
	price := 100
	customer := "customer001"
	filter := invoicemodel.Filter{
		SearchKey:         "mockSearchKey",
		DateFromCreatedAt: &dateInt,
		DateToCreatedAt:   &dateInt,
		MinPrice:          &price,
		MaxPrice:          &price,
		CreatedBy:         &createdBy,
		Customer:          &customer,
	}
	mockProperties := []string{"id"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	listInvoice := []invoicemodel.Invoice{
		{
			Id:                  "123",
			CustomerId:          customer,
			TotalPrice:          100.0,
			AmountReceived:      75,
			AmountPriceUsePoint: 25.0,
			CreatedBy:           createdBy,
			CreatedAt:           &dateTime,
		},
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"customerId",
		"totalPrice",
		"amountReceived",
		"amountPriceUsePoint",
		"createdBy",
		"createdAt",
	})
	for _, invoice := range listInvoice {
		rows.AddRow(
			invoice.Id,
			invoice.CustomerId,
			invoice.TotalPrice,
			invoice.AmountReceived,
			invoice.AmountPriceUsePoint,
			invoice.CreatedBy,
			invoice.CreatedAt,
		)
	}

	queryString :=
		"SELECT `Invoice`.`id`,`Invoice`.`customerId`,`Invoice`.`totalPrice`,`Invoice`.`amountReceived`,`Invoice`.`amountPriceUsePoint`,`Invoice`.`createdBy`,`Invoice`.`createdAt` FROM `Invoice` " +
			"JOIN MUser ON Invoice.createdBy = MUser.id " +
			"JOIN Customer ON Invoice.customerId = Customer.id " +
			"WHERE id LIKE ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"AND totalPrice >= ? " +
			"AND totalPrice <= ? " +
			"AND MUser.Id = ? " +
			"AND Customer.Id = ? " +
			"ORDER BY createdAt desc LIMIT 10"

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listInvoice))
	queryStringCount :=
		"SELECT count(*) FROM `Invoice` " +
			"JOIN MUser ON Invoice.createdBy = MUser.id " +
			"JOIN Customer ON Invoice.customerId = Customer.id " +
			"WHERE id LIKE ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"AND totalPrice >= ? " +
			"AND totalPrice <= ? " +
			"AND MUser.Id = ? " +
			"AND Customer.Id = ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *invoicemodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
		moreKeys                   []string
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
			name:   "List invoice failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						dateTime,
						dateTime,
						price,
						price,
						createdBy,
						customer,
					).
					WillReturnError(mockErr)
			},
			want:    listInvoice,
			wantErr: true,
		},
		{
			name:   "List invoice failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						dateTime,
						dateTime,
						price,
						price,
						createdBy,
						customer,
					).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						dateTime,
						dateTime,
						price,
						price,
						createdBy,
						customer,
					).
					WillReturnError(mockErr)
			},
			want:    listInvoice,
			wantErr: true,
		},
		{
			name:   "List invoice successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						dateTime,
						dateTime,
						price,
						price,
						createdBy,
						customer,
					).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						dateTime,
						dateTime,
						price,
						price,
						createdBy,
						customer,
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

			got, err := s.ListInvoice(
				tt.args.ctx,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListInvoice() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListInvoice() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
