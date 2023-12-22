package supplierstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_ListSupplier(t *testing.T) {
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

	supplier1 := suppliermodel.Supplier{
		Id:    "1",
		Name:  "Supplier1",
		Email: "supplier1@example.com",
		Phone: "123456789",
		Debt:  100.0,
	}
	supplier2 := suppliermodel.Supplier{
		Id:    "2",
		Name:  "Supplier2",
		Email: "supplier2@example.com",
		Phone: "987654321",
		Debt:  200.0,
	}

	minDebt := float32(150.0)
	maxDebt := float32(300.0)
	filterSupplier := &filter.Filter{
		SearchKey: "Supplier",
		MinDebt:   &minDebt,
		MaxDebt:   &maxDebt,
	}

	paging := &common.Paging{
		Page:  2,
		Limit: 10,
	}

	expectedSql := "SELECT * FROM `Supplier` WHERE (id LIKE ? OR name LIKE ? OR email LIKE ? OR phone LIKE ?) AND debt >= ? AND debt <= ? ORDER BY name LIMIT 10 OFFSET 10"
	expectedCountSql := "SELECT count(*) FROM `Supplier` WHERE (id LIKE ? OR name LIKE ? OR email LIKE ? OR phone LIKE ?) AND debt >= ? AND debt <= ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *filter.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []suppliermodel.Supplier
		wantErr bool
	}{
		{
			name: "List suppliers successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filterSupplier,
				propertiesContainSearchKey: []string{"id", "name", "email", "phone"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						"%Supplier%", "%Supplier%", "%Supplier%", "%Supplier%",
						minDebt, maxDebt).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						"%Supplier%", "%Supplier%", "%Supplier%", "%Supplier%",
						minDebt, maxDebt).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "debt"}).
						AddRow(supplier1.Id, supplier1.Name, supplier1.Email, supplier1.Phone, supplier1.Debt).
						AddRow(supplier2.Id, supplier2.Name, supplier2.Email, supplier2.Phone, supplier2.Debt))
			},
			want:    []suppliermodel.Supplier{supplier1, supplier2},
			wantErr: false,
		},
		{
			name: "List suppliers failed because can get total records of suppliers",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filterSupplier,
				propertiesContainSearchKey: []string{"id", "name", "email", "phone"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						"%Supplier%", "%Supplier%", "%Supplier%", "%Supplier%",
						minDebt, maxDebt).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List suppliers failed because can not query supplier records",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filterSupplier,
				propertiesContainSearchKey: []string{"id", "name", "email", "phone"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						"%Supplier%", "%Supplier%", "%Supplier%", "%Supplier%",
						minDebt, maxDebt).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						"%Supplier%", "%Supplier%", "%Supplier%", "%Supplier%",
						minDebt, maxDebt).
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

			got, err := s.ListSupplier(tt.args.ctx, tt.args.filter, tt.args.propertiesContainSearchKey, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListSupplier() got = %v, want %v", got, tt.want)
			}
		})
	}
}
