package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateInfoUserStore struct {
	mock.Mock
}

func (m *mockUpdateInfoUserStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	args := m.Called(ctx, conditions, moreInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *mockUpdateInfoUserStore) UpdateInfoUser(
	ctx context.Context,
	id string,
	data *usermodel.UserUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewUpdateInfoUserRepo(t *testing.T) {
	type args struct {
		userStore UpdateInfoUserStore
	}

	mockUserStore := new(mockUpdateInfoUserStore)

	tests := []struct {
		name string
		args args
		want *updateInfoUserRepo
	}{
		{
			name: "Create object has type UpdateInforUserRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &updateInfoUserRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateInfoUserRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewUpdateInfoUserRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateInfoUserRepo_CheckUserStatusPermission(t *testing.T) {
	type fields struct {
		userStore UpdateInfoUserStore
	}
	type args struct {
		ctx    context.Context
		userId string
	}

	mockUserStore := new(mockUpdateInfoUserStore)
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
			repo := &updateInfoUserRepo{
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

func Test_updateInfoUserRepo_UpdateInfoUser(t *testing.T) {
	type fields struct {
		userStore UpdateInfoUserStore
	}
	type args struct {
		ctx    context.Context
		userId string
		data   *usermodel.UserUpdateInfo
	}

	mockUserStore := new(mockUpdateInfoUserStore)
	mockErr := errors.New(mock.Anything)
	userUpdate := &usermodel.UserUpdateInfo{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "UpdateInfoUser failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: "user123",
				data:   userUpdate,
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdateInfoUser",
						context.Background(),
						"user123",
						userUpdate,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "UpdateInfoUser successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:    context.Background(),
				userId: "user123",
				data:   userUpdate,
			},
			mock: func() {
				mockUserStore.
					On(
						"UpdateInfoUser",
						context.Background(),
						"user123",
						userUpdate,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateInfoUserRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			err := repo.UpdateInfoUser(tt.args.ctx, tt.args.userId, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "UpdateInfoUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateInfoUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
