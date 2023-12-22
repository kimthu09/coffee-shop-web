package customerstore

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_FindCustomer(t *testing.T) {
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

	mockCustomerId := "123"
	mockConditions := map[string]interface{}{"id": mockCustomerId}
	mockCustomer := &customermodel.Customer{
		Id:    mockCustomerId,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "123456789",
		Point: 100.0,
	}

	queryString :=
		"SELECT * FROM `Customer` WHERE `id` = ? ORDER BY `Customer`.`id` LIMIT 1"

	mockErr := errors.New("mock.Anything")

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *customermodel.Customer
		wantErr bool
	}{
		{
			name:   "Find customer failed because of a database error",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(queryString).
					WithArgs(mockCustomerId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Find customer failed because the customer does not exist",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(queryString).
					WithArgs(mockCustomerId).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Find customer successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "point"}).
					AddRow(mockCustomer.Id, mockCustomer.Name, mockCustomer.Email, mockCustomer.Phone, mockCustomer.Point)

				sqlDBMock.
					ExpectQuery(queryString).
					WithArgs(mockCustomerId).
					WillReturnRows(rows)
			},
			want:    mockCustomer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindCustomer(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindCustomer() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindCustomer() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
