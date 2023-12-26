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

func Test_sqlStore_UpdateCustomerInfo(t *testing.T) {
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

	customerId := mock.Anything
	updateData := &customermodel.CustomerUpdateInfo{
		Name:  new(string),
		Email: new(string),
		Phone: new(string),
	}

	expectedSql := "UPDATE `Customer` SET `name`=?,`email`=?,`phone`=? WHERE id = ?"

	mockErr := errors.New(mock.Anything)
	mockErrPhone := &common.GormErr{
		Number:  1062,
		Message: "phone",
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *customermodel.CustomerUpdateInfo
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
			name: "Update customer info successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(updateData.Name, updateData.Email, updateData.Phone, customerId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Update customer info failed because can not save to database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(updateData.Name, updateData.Email, updateData.Phone, customerId).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Update customer info failed because duplicate phone",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   customerId,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(updateData.Name, updateData.Email, updateData.Phone, customerId).
					WillReturnError(mockErrPhone)
				sqlDBMock.ExpectRollback()
			},
			want:    customermodel.ErrCustomerPhoneDuplicate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.UpdateCustomerInfo(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateCustomerInfo() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "UpdateCustomerInfo() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "UpdateCustomerInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
