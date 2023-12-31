package userstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateUser(t *testing.T) {
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

	userId := mock.Anything
	userName := mock.Anything
	userEmail := mock.Anything
	userRoleId := mock.Anything
	userSalt := mock.Anything
	userPass := mock.Anything
	userCreate := usermodel.UserCreate{
		Id:       userId,
		Name:     userName,
		Email:    userEmail,
		Password: userPass,
		Salt:     userSalt,
		RoleId:   userRoleId,
	}
	mockErr := errors.New("something went wrong with the database")
	mockErrEmail := &common.GormErr{
		Number:  1062,
		Message: "email",
	}

	expectedSql := "INSERT INTO `MUser` (`id`,`name`,`email`,`password`,`salt`,`roleId`) VALUES (?,?,?,?,?,?)"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *usermodel.UserCreate
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
			name: "Create user in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &userCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						userCreate.Id,
						userCreate.Name,
						userCreate.Email,
						userCreate.Password,
						userCreate.Salt,
						userCreate.RoleId,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Create user in database failed because of database error",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &userCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						userCreate.Id,
						userCreate.Name,
						userCreate.Email,
						userCreate.Password,
						userCreate.Salt,
						userCreate.RoleId,
					).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Create user in database failed because of duplicate email",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &userCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						userCreate.Id,
						userCreate.Name,
						userCreate.Email,
						userCreate.Password,
						userCreate.Salt,
						userCreate.RoleId,
					).
					WillReturnError(mockErrEmail)
				sqlDBMock.ExpectRollback()
			},
			want:    usermodel.ErrUserEmailDuplicated,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateUser(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateUser() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
