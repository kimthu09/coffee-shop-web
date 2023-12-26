package rolestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_UpdateRole(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *rolemodel.RoleUpdate
	}

	roleId := "Role001"
	name := "New name"
	mockData := &rolemodel.RoleUpdate{
		Name: &name,
	}

	expectedQuery := "UPDATE `Role` SET `name`=? WHERE id = ?"
	mockErr := errors.New(mock.Anything)
	mockErrName := &common.GormErr{
		Number:  1062,
		Message: "name",
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
			name: "Update role successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   roleId,
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.Name, roleId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Update role failed because of database error",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   roleId,
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.Name, roleId).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Update role failed because of duplicate name",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   roleId,
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockData.Name, roleId).
					WillReturnError(mockErrName)
				sqlDBMock.ExpectRollback()
			},
			want:    rolemodel.ErrRoleNameDuplicate,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.UpdateRole(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "UpdateRole() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "UpdateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
