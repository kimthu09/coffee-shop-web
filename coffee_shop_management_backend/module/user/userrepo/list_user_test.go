package userrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListUserStore struct {
	mock.Mock
}

func (m *mockListUserStore) ListUser(
	ctx context.Context,
	userSearch string,
	filter *usermodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.User, error) {
	args := m.Called(ctx, userSearch, filter, propertiesContainSearchKey, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]usermodel.User), args.Error(1)
}

func TestNewListUserRepo(t *testing.T) {
	type args struct {
		store ListUserStore
	}

	mockUserStore := new(mockListUserStore)

	tests := []struct {
		name string
		args args
		want *listUserRepo
	}{
		{
			name: "Create object has type ListUserRepo",
			args: args{
				store: mockUserStore,
			},
			want: &listUserRepo{
				store: mockUserStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListUserRepo(
				tt.args.store,
			)

			assert.Equal(t, tt.want, got, "NewListUserRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listUserRepo_ListUser(t *testing.T) {
	type fields struct {
		store ListUserStore
	}
	type args struct {
		ctx        context.Context
		userSearch string
		filter     *usermodel.Filter
		paging     *common.Paging
	}

	mockUserStore := new(mockListUserStore)
	userData := make([]usermodel.User, 0)
	filter := usermodel.Filter{}
	moreKeys := []string{"Role"}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []usermodel.User
		wantErr bool
	}{
		{
			name: "ListUser failed",
			fields: fields{
				store: mockUserStore,
			},
			args: args{
				ctx:        context.Background(),
				userSearch: "User001",
				filter:     &usermodel.Filter{},
				paging:     &common.Paging{},
			},
			mock: func() {
				mockUserStore.
					On(
						"ListUser",
						context.Background(),
						"User001",
						&filter,
						[]string{"id", "name", "email", "phone", "address"},
						&common.Paging{},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    userData,
			wantErr: true,
		},
		{
			name: "ListUser successfully",
			fields: fields{
				store: mockUserStore,
			},
			args: args{
				ctx:        context.Background(),
				userSearch: "User001",
				filter:     &usermodel.Filter{},
				paging:     &common.Paging{},
			},
			mock: func() {
				mockUserStore.
					On(
						"ListUser",
						context.Background(),
						"User001",
						&filter,
						[]string{"id", "name", "email", "phone", "address"},
						&common.Paging{},
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
			repo := &listUserRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListUser(
				tt.args.ctx, tt.args.userSearch, tt.args.filter, tt.args.paging,
			)
			if tt.wantErr {
				assert.NotNil(t, err, "ListUser() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListUser() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListUser() got = %v, want %v", got, tt.want)

			}
		})
	}
}
