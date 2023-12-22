package supplierstore

import (
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

func Test_sqlStore_FindSupplier(t *testing.T) {
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

	supplierId := mock.Anything
	supplierName := mock.Anything
	supplierEmail := mock.Anything
	supplierPhone := mock.Anything
	supplierDebt := -100000
	supplier := suppliermodel.Supplier{
		Id:    supplierId,
		Name:  supplierName,
		Email: supplierEmail,
		Phone: supplierPhone,
		Debt:  supplierDebt,
	}
	conditions := map[string]interface{}{"id": supplierId}
	mockErr := errors.New("something went wrong with the database")

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	expectedSql := "SELECT * FROM `Supplier` WHERE `id` = ? ORDER BY `Supplier`.`id` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *suppliermodel.Supplier
		wantErr bool
	}{
		{
			name: "Find supplier in database successfully",
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
					WithArgs(supplierId).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "email", "phone", "debt"}).
							AddRow(supplier.Id, supplier.Name, supplier.Email, supplier.Phone, supplier.Debt))
			},
			want:    &supplier,
			wantErr: false,
		},
		{
			name: "Find supplier in database failed something went wrong with the database",
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
					WithArgs(supplierId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find supplier in database record not found",
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
					WithArgs(supplierId).
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

			got, err := s.FindSupplier(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindSupplier() got = %v, want %v", got, tt.want)
			}
		})
	}
}
