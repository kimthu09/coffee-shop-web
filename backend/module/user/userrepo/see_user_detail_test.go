package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockFindUserStore struct {
	mock.Mock
}

func (m *mockFindUserStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	args := m.Called(ctx, conditions, moreInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func TestNewSeeUserDetailRepo(t *testing.T) {
	type args struct {
		userStore FindUserStore
	}
	mockUserStore := new(mockFindUserStore)

	tests := []struct {
		name string
		args args
		want *seeUserDetailRepo
	}{
		{
			name: "Create object has type SeeUserDetailRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &seeUserDetailRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeUserDetailRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewSeeUserDetailRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeUserDetailRepo_SeeUserDetail(t *testing.T) {
	type fields struct {
		userStore FindUserStore
	}
	type args struct {
		ctx    context.Context
		userId string
	}

	mockUserStore := new(mockFindUserStore)
	userData := &usermodel.User{}
	mockErr := errors.New(mock.Anything)
	moreKeys := []string{"Role.RoleFeatures"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *usermodel.User
		wantErr bool
	}{
		{
			name: "SeeUserDetail failed",
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
			want:    nil,
			wantErr: true,
		},
		{
			name: "SeeUserDetail successfully",
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
			repo := &seeUserDetailRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			got, err := repo.SeeUserDetail(tt.args.ctx, tt.args.userId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeUserDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeUserDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeUserDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
