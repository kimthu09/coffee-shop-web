package userstore

import (
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

func Test_sqlStore_FindUser(t *testing.T) {
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

	userId := mock.Anything
	userName := mock.Anything
	userEmail := mock.Anything
	userRoleId := mock.Anything
	userSalt := mock.Anything
	userPass := mock.Anything
	user := usermodel.User{
		Id:       userId,
		Name:     userName,
		Email:    userEmail,
		Password: userPass,
		Salt:     userSalt,
		RoleId:   userRoleId,
		IsActive: true,
	}
	conditions := map[string]interface{}{"id": userId}
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	expectedSql := "SELECT * FROM `MUser` WHERE `id` = ? ORDER BY `MUser`.`id` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *usermodel.User
		wantErr bool
	}{
		{
			name: "Find user in database successfully",
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
					WithArgs(userId).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name", "email", "password", "salt", "roleId", "isActive"}).
							AddRow(user.Id, user.Name, user.Email, user.Password, user.Salt, user.RoleId, user.IsActive))
			},
			want:    &user,
			wantErr: false,
		},
		{
			name: "Find user in database failed something went wrong with the database",
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
					WithArgs(userId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find user in database record not found",
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
					WithArgs(userId).
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

			got, err := s.FindUser(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindUser() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
