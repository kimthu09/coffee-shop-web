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

func Test_sqlStore_CreateCustomer(t *testing.T) {
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

	id := "123"
	mockCustomer := &customermodel.CustomerCreate{
		Id:    &id,
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Phone: "123456789",
	}

	queryString :=
		"INSERT INTO `Customer` (`id`,`name`,`email`,`phone`) VALUES (?,?,?,?)"

	mockErr := errors.New(mock.Anything)
	mockErrPhone := &common.GormErr{
		Number:  1062,
		Message: "phone",
	}
	mockErrPRIMARY := &common.GormErr{
		Number:  1062,
		Message: "PRIMARY",
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *customermodel.CustomerCreate
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
			name:   "Create customer successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: mockCustomer,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(queryString).
					WithArgs(
						mockCustomer.Id,
						mockCustomer.Name,
						mockCustomer.Email,
						mockCustomer.Phone,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:   "Create customer failed because of something went wrong in db",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: mockCustomer,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(queryString).
					WithArgs(
						mockCustomer.Id,
						mockCustomer.Name,
						mockCustomer.Email,
						mockCustomer.Phone,
					).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name:   "Create customer failed because of duplicate phone",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: mockCustomer,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(queryString).
					WithArgs(
						mockCustomer.Id,
						mockCustomer.Name,
						mockCustomer.Email,
						mockCustomer.Phone,
					).
					WillReturnError(mockErrPhone)
				sqlDBMock.ExpectRollback()
			},
			want:    customermodel.ErrCustomerPhoneDuplicate,
			wantErr: true,
		},
		{
			name:   "Create customer failed because of duplicate id",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: mockCustomer,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(queryString).
					WithArgs(
						mockCustomer.Id,
						mockCustomer.Name,
						mockCustomer.Email,
						mockCustomer.Phone,
					).
					WillReturnError(mockErrPRIMARY)
				sqlDBMock.ExpectRollback()
			},
			want:    customermodel.ErrCustomerIdDuplicate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateCustomer(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateCustomer() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateCustomer() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateCustomer() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
