package customerstore

import (
	"coffee_shop_management_backend/common"
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

func Test_sqlStore_ListCustomer(t *testing.T) {
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

	customer1 := customermodel.Customer{
		Id:    "1",
		Name:  "Customer1",
		Email: "customer1@example.com",
		Phone: "123456789",
		Point: 100.0,
	}
	customer2 := customermodel.Customer{
		Id:    "2",
		Name:  "Customer2",
		Email: "customre2@example.com",
		Phone: "987654321",
		Point: 200.0,
	}

	minPoint := float32(50.0)
	maxPoint := float32(300.0)
	filterCustomer := &customermodel.Filter{
		SearchKey: "Customer",
		MinPoint:  &minPoint,
		MaxPoint:  &maxPoint,
	}

	paging := &common.Paging{
		Page:  2,
		Limit: 10,
	}

	expectedSql := "SELECT * FROM `Customer` WHERE (id LIKE ? OR name LIKE ? OR email LIKE ? OR phone LIKE ?) AND point >= ? AND point <= ? ORDER BY name LIMIT 10 OFFSET 10"
	expectedCountSql := "SELECT count(*) FROM `Customer` WHERE (id LIKE ? OR name LIKE ? OR email LIKE ? OR phone LIKE ?) AND point >= ? AND point <= ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *customermodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []customermodel.Customer
		wantErr bool
	}{
		{
			name: "List customers successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filterCustomer,
				propertiesContainSearchKey: []string{"id", "name", "email", "phone"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						"%Customer%", "%Customer%", "%Customer%", "%Customer%",
						minPoint, maxPoint).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						"%Customer%", "%Customer%", "%Customer%", "%Customer%",
						minPoint, maxPoint).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "point"}).
						AddRow(customer1.Id, customer1.Name, customer1.Email, customer1.Phone, customer1.Point).
						AddRow(customer2.Id, customer2.Name, customer2.Email, customer2.Phone, customer2.Point))
			},
			want:    []customermodel.Customer{customer1, customer2},
			wantErr: false,
		},
		{
			name: "List customers failed because can get total records of customer",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filterCustomer,
				propertiesContainSearchKey: []string{"id", "name", "email", "phone"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						"%Customer%", "%Customer%", "%Customer%", "%Customer%",
						minPoint, maxPoint).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List customers failed because can not query customer records",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                        context.Background(),
				filter:                     filterCustomer,
				propertiesContainSearchKey: []string{"id", "name", "email", "phone"},
				paging:                     paging,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedCountSql).
					WithArgs(
						"%Customer%", "%Customer%", "%Customer%", "%Customer%",
						minPoint, maxPoint).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						"%Customer%", "%Customer%", "%Customer%", "%Customer%",
						minPoint, maxPoint).
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

			got, err := s.ListCustomer(tt.args.ctx, tt.args.filter, tt.args.propertiesContainSearchKey, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListCustomer() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListCustomer() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}
