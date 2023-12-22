package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockResetPasswordStore struct {
	mock.Mock
}

func (m *mockResetPasswordStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {
	args := m.Called(ctx, conditions, moreInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *mockResetPasswordStore) UpdatePasswordUser(
	ctx context.Context,
	id string,
	password string,
) error {
	args := m.Called(ctx, id, password)
	return args.Error(0)
}

func TestNewResetPasswordRepo(t *testing.T) {
	type args struct {
		userStore ResetPasswordStore
	}

	mockUserStore := new(mockResetPasswordStore)

	tests := []struct {
		name string
		args args
		want *resetPasswordRepo
	}{
		{
			name: "Create object has type ResetPasswordRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &resetPasswordRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewResetPasswordRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewResetPasswordRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_resetPasswordRepo_CheckUserStatusPermission(t *testing.T) {
	type fields struct {
		userStore ResetPasswordStore
	}
	type args struct {
		ctx    context.Context
		userId string
	}

	mockUserStore := new(mockResetPasswordStore)
	inactiveUserData := &usermodel.User{
		IsActive: false,
	}
	userData := &usermodel.User{
		IsActive: true,
	}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *usermodel.User
		wantErr bool
	}{
		{
			name: "CheckUserStatusPermission failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: "user123",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": "user123"},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    userData,
			wantErr: true,
		},
		{
			name: "CheckUserStatusPermission inactive user",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: "user123",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": "user123"},
						moreKeys).
					Return(inactiveUserData, nil).
					Once()
			},
			want:    inactiveUserData,
			wantErr: true,
		},
		{
			name: "CheckUserStatusPermission successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: "user123",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": "user123"},
						moreKeys).
					Return(userData, nil).
					Once()
			},
			want:    userData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &resetPasswordRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			err := repo.CheckUserStatusPermission(tt.args.ctx, tt.args.userId)
			if tt.wantErr {
				assert.NotNil(t, err, "CheckUserStatusPermission() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CheckUserStatusPermission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_resetPasswordRepo_GetUser(t *testing.T) {
	type fields struct {
		userStore ResetPasswordStore
	}
	type args struct {
		ctx context.Context
		id  string
	}

	mockUserStore := new(mockResetPasswordStore)
	userData := &usermodel.User{}
	mockErr := errors.New(mock.Anything)
	var moreKeys []string

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *usermodel.User
		wantErr bool
	}{
		{
			name: "GetUser failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx: context.Background(),
				id:  "user123",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": "user123"},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    userData,
			wantErr: true,
		},
		{
			name: "GetUser successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx: context.Background(),
				id:  "user123",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": "user123"},
						moreKeys).
					Return(userData, nil).
					Once()
			},
			want:    userData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &resetPasswordRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			got, err := repo.GetUser(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.NotNil(t, err, "GetUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetUser() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetUser() got = %v, want %v", got, tt.want)

			}
		})
	}
}

func Test_resetPasswordRepo_UpdateUserPassword(t *testing.T) {
	type fields struct {
		userStore ResetPasswordStore
	}
	type args struct {
		ctx  context.Context
		id   string
		pass string
	}

	mockUserStore := new(mockResetPasswordStore)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "UpdateUserPassword failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:  context.Background(),
				id:   "user123",
				pass: "newPass123",
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdatePasswordUser",
						context.Background(),
						"user123",
						"newPass123").
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "UpdateUserPassword successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:  context.Background(),
				id:   "user123",
				pass: "newPass123",
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdatePasswordUser",
						context.Background(),
						"user123",
						"newPass123").
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &resetPasswordRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			err := repo.UpdateUserPassword(tt.args.ctx, tt.args.id, tt.args.pass)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateUserPassword() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateUserPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
