package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdatePasswordStore struct {
	mock.Mock
}

func (m *mockUpdatePasswordStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	args := m.Called(ctx, conditions, moreInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *mockUpdatePasswordStore) UpdatePasswordUser(
	ctx context.Context,
	id string,
	password string) error {
	args := m.Called(ctx, id, password)
	return args.Error(0)
}

func TestNewUpdatePasswordRepo(t *testing.T) {
	type args struct {
		userStore ResetPasswordStore
	}

	mockUserStore := new(mockUpdatePasswordStore)

	tests := []struct {
		name string
		args args
		want *updatePasswordRepo
	}{
		{
			name: "Create object has type UpdatePasswordRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &updatePasswordRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdatePasswordRepo(tt.args.userStore)

			assert.Equal(t, tt.want, got, "NewUpdatePasswordRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updatePasswordRepo_GetUser(t *testing.T) {
	type fields struct {
		userStore ResetPasswordStore
	}
	type args struct {
		ctx    context.Context
		userId string
	}

	mockUserStore := new(mockUpdatePasswordStore)
	mockErr := errors.New("mock error")
	userId := "user123"
	var moreKeys []string
	user := &usermodel.User{}

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
				ctx:    context.Background(),
				userId: userId,
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": userId},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetUser successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"id": userId},
						moreKeys,
					).
					Return(user, nil).
					Once()
			},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updatePasswordRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			got, err := repo.GetUser(tt.args.ctx, tt.args.userId)
			if tt.wantErr {
				assert.NotNil(t, err, "GetUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetUser() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updatePasswordRepo_UpdateUserPassword(t *testing.T) {
	type fields struct {
		userStore ResetPasswordStore
	}
	type args struct {
		ctx  context.Context
		id   string
		pass string
	}

	mockUserStore := new(mockUpdatePasswordStore)
	mockErr := errors.New("mock error")

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
				pass: "newpass",
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdatePasswordUser",
						context.Background(),
						"user123",
						"newpass",
					).
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
				pass: "newpass",
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdatePasswordUser",
						context.Background(),
						"user123",
						"newpass",
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updatePasswordRepo{
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
