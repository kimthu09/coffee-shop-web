package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeStatusUserStore struct {
	mock.Mock
}

func (m *mockChangeStatusUserStore) UpdateStatusUser(ctx context.Context, data *usermodel.UserUpdateStatus) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewChangeStatusUserRepo(t *testing.T) {
	type args struct {
		userStore ChangeStatusUserStore
	}

	mockUserStore := new(mockChangeStatusUserStore)

	tests := []struct {
		name string
		args args
		want *changeStatusUserRepo
	}{
		{
			name: "Create object has type ChangeRoleUserRepo",
			args: args{
				userStore: mockUserStore,
			},
			want: &changeStatusUserRepo{
				userStore: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusUserRepo(
				tt.args.userStore,
			)

			assert.Equal(t, tt.want, got, "NewChangeStatusUserRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusUserRepo_UpdateStatusUser(t *testing.T) {
	type fields struct {
		userStore ChangeStatusUserStore
	}
	type args struct {
		ctx  context.Context
		data *usermodel.UserUpdateStatus
	}

	mockUserStore := new(mockChangeStatusUserStore)
	active := true
	data := &usermodel.UserUpdateStatus{IsActive: &active}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "UpdateStatusUser failed",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:  context.Background(),
				data: data,
			},
			mock: func() {
				mockUserStore.On("UpdateStatusUser", context.Background(), data).Return(errors.New("update status error")).Once()
			},
			wantErr: true,
		},
		{
			name: "UpdateStatusUser successfully",
			fields: fields{
				userStore: mockUserStore,
			},
			args: args{
				ctx:  context.Background(),
				data: data,
			},
			mock: func() {
				mockUserStore.On("UpdateStatusUser", context.Background(), data).Return(nil).Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusUserRepo{
				userStore: tt.fields.userStore,
			}

			tt.mock()

			err := repo.UpdateStatusUser(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "UpdateStatusUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateStatusUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
