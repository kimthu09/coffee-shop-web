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

func Test_sqlStore_CreateRole(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *rolemodel.RoleCreate
	}

	roleCreate := rolemodel.RoleCreate{
		Id:   mock.Anything,
		Name: mock.Anything,
	}
	expectedSql := "INSERT INTO `Role` (`id`,`name`) VALUES (?,?)"
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
			name: "Create role in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &roleCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						roleCreate.Id,
						roleCreate.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Create role failed because of database error",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &roleCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						roleCreate.Id,
						roleCreate.Name).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Create role failed because of duplicate name",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &roleCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						roleCreate.Id,
						roleCreate.Name).
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

			err := s.CreateRole(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateRole() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateRole() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
