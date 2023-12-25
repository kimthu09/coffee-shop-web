package supplierstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateSupplier(t *testing.T) {
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

	supplierId := mock.Anything
	supplierName := mock.Anything
	supplierEmail := mock.Anything
	supplierPhone := mock.Anything
	supplierDebt := -1000
	supplierCreate := suppliermodel.SupplierCreate{
		Id:    &supplierId,
		Name:  supplierName,
		Email: supplierEmail,
		Phone: supplierPhone,
		Debt:  supplierDebt,
	}
	mockErr := errors.New("something went wrong with the database")
	mockErrPhone := &common.GormErr{
		Number:  1062,
		Message: "phone",
	}
	mockErrPRIMARY := &common.GormErr{
		Number:  1062,
		Message: "PRIMARY",
	}

	expectedSql := "INSERT INTO `Supplier` (`id`,`name`,`email`,`phone`,`debt`) VALUES (?,?,?,?,?)"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *suppliermodel.SupplierCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    error
		wantErr bool
	}{
		{
			name: "Create supplier in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						supplierCreate.Id,
						supplierCreate.Name,
						supplierCreate.Email,
						supplierCreate.Phone,
						supplierCreate.Debt,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Create supplier in database failed because of database error",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						supplierCreate.Id,
						supplierCreate.Name,
						supplierCreate.Email,
						supplierCreate.Phone,
						supplierCreate.Debt,
					).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Create supplier in database failed because of duplicate phone",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						supplierCreate.Id,
						supplierCreate.Name,
						supplierCreate.Email,
						supplierCreate.Phone,
						supplierCreate.Debt,
					).
					WillReturnError(mockErrPhone)
				sqlDBMock.ExpectRollback()
			},
			want:    suppliermodel.ErrSupplierPhoneDuplicate,
			wantErr: true,
		},
		{
			name: "Create supplier in database failed because of duplicate id",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &supplierCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						supplierCreate.Id,
						supplierCreate.Name,
						supplierCreate.Email,
						supplierCreate.Phone,
						supplierCreate.Debt,
					).
					WillReturnError(mockErrPRIMARY)
				sqlDBMock.ExpectRollback()
			},
			want:    suppliermodel.ErrSupplierIdDuplicate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateSupplier(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateSupplier() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
