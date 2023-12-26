package supplierdebtstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
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

func Test_sqlStore_ListSupplierDebtBySupplier(t *testing.T) {
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

	supplierId := "supplier123"
	date := int64(1639560000)
	filterSupplierDebt := &filter.SupplierDebtFilter{
		DateFrom: &date,
		DateTo:   &date,
	}
	paging := &common.Paging{
		Page:  2,
		Limit: 10,
	}

	expectedSql :=
		"SELECT * FROM `SupplierDebt` WHERE supplierId = ? AND createdAt >= ? AND createdAt <= ? ORDER BY createdAt desc LIMIT 10 OFFSET 10"
	expectedCountSql := "SELECT count(*) FROM `SupplierDebt` WHERE supplierId = ? AND createdAt >= ? AND createdAt <= ?"
	pay := enum.Pay
	dateFrom := time.Unix(*filterSupplierDebt.DateFrom, 0)
	dateTo := time.Unix(*filterSupplierDebt.DateTo, 0)

	mockErr := errors.New(mock.Anything)
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		supplierId string
		filter     *filter.SupplierDebtFilter
		paging     *common.Paging
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []supplierdebtmodel.SupplierDebt
		wantErr bool
	}{
		{
			name: "List supplier debts by supplier successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     filterSupplierDebt,
				paging:     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						supplierId,
						&dateFrom,
						&dateTo).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(11))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						supplierId,
						&dateFrom,
						&dateTo).
					WillReturnRows(
						sqlmock.NewRows(
							[]string{"id", "supplierId", "amount",
								"amountLeft", "type", "createdBy", "createdAt"},
						).AddRow(
							"1", supplierId, 100,
							200, pay, "admin", nil))
			},
			want: []supplierdebtmodel.SupplierDebt{
				{
					Id:         "1",
					SupplierId: supplierId,
					Amount:     100,
					AmountLeft: 200,
					DebtType:   &pay,
					CreatedBy:  "admin"},
			},
			wantErr: false,
		},
		{
			name: "List supplier debts by supplier failed " +
				"because can get total records of supplier debts",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     filterSupplierDebt,
				paging:     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						supplierId,
						&dateFrom,
						&dateTo).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List supplier debts by supplier failed " +
				"because can not query supplier debt records",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
				filter:     filterSupplierDebt,
				paging:     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						supplierId,
						&dateFrom,
						&dateTo).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(11))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						supplierId,
						&dateFrom,
						&dateTo).
					WillReturnError(mockErr)
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

			got, err := s.ListSupplierDebt(
				tt.args.ctx,
				tt.args.supplierId,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListSupplierDebtBySupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListSupplierDebtBySupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListSupplierDebtBySupplier() got = %v, want %v", got, tt.want)
			}
		})
	}
}
