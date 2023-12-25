package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateUserStore struct {
	mock.Mock
}

func (m *mockCreateUserStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateUserRepo(t *testing.T) {
	type args struct {
		userStore CreateUserStore
	}

	mockUserStore := new(mockCreateUserStore)

	tests := []struct {
		name string
		args args
		want *createUserRepo
	}{
		{
			name: "Create object has type CreateUserRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &createUserRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateUserRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewCreateUserRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createUserRepo_CreateUser(t *testing.T) {
	type fields struct {
		userStore CreateUserStore
	}
	type args struct {
		ctx  context.Context
		data *usermodel.UserCreate
	}

	mockUserStore := new(mockCreateUserStore)
	data := &usermodel.UserCreate{}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "CreateUser failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:  context.Background(),
				data: data,
			},
			mock: func() {
				mockUserStore.On("CreateUser", context.Background(), data).Return(errors.New("create user error")).Once()
			},
			wantErr: true,
		},
		{
			name: "CreateUser successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:  context.Background(),
				data: data,
			},
			mock: func() {
				mockUserStore.On("CreateUser", context.Background(), data).Return(nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createUserRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			err := repo.CreateUser(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
