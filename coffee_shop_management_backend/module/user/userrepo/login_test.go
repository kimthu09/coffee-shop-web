package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockLoginStore struct {
	mock.Mock
}

func (m *mockLoginStore) FindUser(
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

func TestNewLoginRepo(t *testing.T) {
	type args struct {
		userStore LoginStore
	}

	mockUserStore := new(mockLoginStore)

	tests := []struct {
		name string
		args args
		want *loginRepo
	}{
		{
			name: "Create object has type LoginRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &loginRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLoginRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewLoginRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_loginRepo_FindUserByEmail(t *testing.T) {
	type fields struct {
		userStore LoginStore
	}
	type args struct {
		ctx   context.Context
		email string
	}

	mockUserStore := new(mockLoginStore)
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
			name: "FindUserByEmail failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:   context.Background(),
				email: "test@example.com",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"email": "test@example.com"},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FindUserByEmail successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:   context.Background(),
				email: "test@example.com",
			},
			mock: func() {
				mockUserStore.
					On(
						"FindUser",
						context.Background(),
						map[string]interface{}{"email": "test@example.com"},
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
			repo := &loginRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			got, err := repo.FindUserByEmail(tt.args.ctx, tt.args.email)
			if tt.wantErr {
				assert.NotNil(t, err, "FindUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindUserByEmail() got = %v, want %v", got, tt.want)

			}
		})
	}
}
