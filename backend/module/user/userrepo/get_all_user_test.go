package userrepo

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockGetAllUserStore struct {
	mock.Mock
}

func (m *mockGetAllUserStore) GetAllUser(ctx context.Context, moreKeys ...string) ([]usermodel.SimpleUser, error) {
	args := m.Called(ctx, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]usermodel.SimpleUser), args.Error(1)
}

func TestNewGetAllUserRepo(t *testing.T) {
	type args struct {
		store GetAllUserStore
	}

	mockUserStore := new(mockGetAllUserStore)

	tests := []struct {
		name string
		args args
		want *getAllUserRepo
	}{
		{
			name: "Create object has type GetAllUserRepo",
			args: args{
				store: mockUserStore,
			},
			want: &getAllUserRepo{
				store: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetAllUserRepo(
				tt.args.store,
			)

			assert.Equal(t, tt.want, got, "NewGetAllUserRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_getAllUserRepo_GetAllUser(t *testing.T) {
	type fields struct {
		store GetAllUserStore
	}
	type args struct {
		ctx context.Context
	}

	mockUserStore := new(mockGetAllUserStore)
	userData := make([]usermodel.SimpleUser, 0)
	mockErr := errors.New(mock.Anything)
	var moreKeys []string

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []usermodel.SimpleUser
		wantErr bool
	}{
		{
			name: "Get all user failed",
			fields: fields{
				store: mockUserStore,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockUserStore.
					On(
						"GetAllUser",
						context.Background(),
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get all user successfully",
			fields: fields{
				store: mockUserStore,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockUserStore.
					On(
						"GetAllUser",
						context.Background(),
						moreKeys,
					).
					Return(userData, nil).
					Once()
			},
			want:    userData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &getAllUserRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.GetAllUser(tt.args.ctx)
			if tt.wantErr {
				assert.NotNil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetAllSupplier() got = %v, want %v", got, tt.want)

			}
		})
	}
}
