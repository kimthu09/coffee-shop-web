package supplierdebtstore

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateSupplierDebt(t *testing.T) {
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

	payDebtType := enum.Pay
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         "1",
		SupplierId: "supplier123",
		Amount:     100,
		AmountLeft: -200,
		DebtType:   &payDebtType,
		CreatedBy:  "admin",
	}

	expectedSql := "INSERT INTO `SupplierDebt` (`id`,`supplierId`,`amount`,`amountLeft`,`type`,`createdBy`) VALUES (?,?,?,?,?,?)"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *supplierdebtmodel.SupplierDebtCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create supplier debt successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierDebtCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(
						supplierDebtCreate.Id,
						supplierDebtCreate.SupplierId,
						supplierDebtCreate.Amount,
						supplierDebtCreate.AmountLeft,
						supplierDebtCreate.DebtType,
						supplierDebtCreate.CreatedBy).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create supplier debt failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierDebtCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(
						supplierDebtCreate.Id,
						supplierDebtCreate.SupplierId,
						supplierDebtCreate.Amount,
						supplierDebtCreate.AmountLeft,
						supplierDebtCreate.DebtType,
						supplierDebtCreate.CreatedBy).
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

			err := s.CreateSupplierDebt(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
