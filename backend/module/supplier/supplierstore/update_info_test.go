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

func Test_sqlStore_UpdateSupplierInfo(t *testing.T) {
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

	supplierID := mock.Anything
	name := "Updated name"
	phone := "Updated phone"
	email := "Updated email"
	updateData := &suppliermodel.SupplierUpdateInfo{
		Name:  &name,
		Email: &email,
		Phone: &phone,
	}

	expectedSql := "UPDATE `Supplier` SET `name`=?,`email`=?,`phone`=? WHERE id = ?"

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
		data *suppliermodel.SupplierUpdateInfo
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
			name: "Update supplier info successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierID,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(updateData.Name, updateData.Email, updateData.Phone, supplierID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Update supplier info failed because of database error",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierID,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(updateData.Name, updateData.Email, updateData.Phone, supplierID).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Update supplier info failed because of duplicate phone",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   supplierID,
				data: updateData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(updateData.Name, updateData.Email, updateData.Phone, supplierID).
					WillReturnError(mockErrPhone)
				sqlDBMock.ExpectRollback()
			},
			want:    suppliermodel.ErrSupplierPhoneDuplicate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.UpdateSupplierInfo(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateSupplierInfo() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "UpdateSupplierInfo() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "UpdateSupplierInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
