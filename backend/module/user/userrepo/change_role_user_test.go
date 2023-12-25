package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeRoleUserStore struct {
	mock.Mock
}

func (m *mockChangeRoleUserStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	args := m.Called(ctx, conditions, moreInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *mockChangeRoleUserStore) UpdateRoleUser(
	ctx context.Context,
	id string,
	data *usermodel.UserUpdateRole) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewChangeRoleUserRepo(t *testing.T) {
	type args struct {
		userStore ChangeRoleUserStore
	}

	mockUserStore := new(mockChangeRoleUserStore)

	tests := []struct {
		name string
		args args
		want *changeRoleUserRepo
	}{
		{
			name: "Create object has type ChangeRoleUserRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &changeRoleUserRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeRoleUserRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewChangeRoleUserRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeRoleUserRepo_CheckUserStatusPermission(t *testing.T) {
	type fields struct {
		userStore ChangeRoleUserStore
	}
	type args struct {
		ctx    context.Context
		userId string
	}

	mockUserStore := new(mockChangeRoleUserStore)
	mockErr := errors.New(mock.Anything)
	userId := "123"
	inActiveUser := usermodel.User{
		IsActive: false,
	}
	activeUser := usermodel.User{
		IsActive: true,
	}
	var moreKeys []string

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "CheckUserStatusPermission failed due to user store error",
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
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "CheckUserStatusPermission failed due to inactive user",
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
						moreKeys).
					Return(&inActiveUser, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "CheckUserStatusPermission successfully",
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
						moreKeys).
					Return(&activeUser, nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeRoleUserRepo{
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

func Test_changeRoleUserRepo_UpdateRoleUser(t *testing.T) {
	type fields struct {
		userStore ChangeRoleUserStore
	}
	type args struct {
		ctx    context.Context
		userId string
		data   *usermodel.UserUpdateRole
	}

	userId := "User001"
	mockUserStore := new(mockChangeRoleUserStore)
	userUpdateRole := usermodel.UserUpdateRole{RoleId: "Admin"}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "UpdateRoleUser failed due to user store error",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
				data:   &userUpdateRole,
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdateRoleUser",
						context.Background(),
						userId,
						&userUpdateRole).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "UpdateRoleUser successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: userId,
				data:   &userUpdateRole,
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdateRoleUser",
						context.Background(),
						userId,
						&userUpdateRole).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeRoleUserRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			err := repo.UpdateRoleUser(tt.args.ctx, tt.args.userId, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRoleUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateRoleUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
