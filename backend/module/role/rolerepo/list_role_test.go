package rolerepo

import (
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListRoleStore struct {
	mock.Mock
}

func (m *mockListRoleStore) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]rolemodel.SimpleRole), args.Error(1)
}

func TestNewListRoleRepo(t *testing.T) {
	type args struct {
		store ListRoleStore
	}

	mockRole := new(mockListRoleStore)

	tests := []struct {
		name string
		args args
		want *listRoleRepo
	}{
		{
			name: "Create object has type CreateRoleRepo",
			args: args{
				store: mockRole,
			},
			want: &listRoleRepo{
				store: mockRole,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListRoleRepo(
				tt.args.store)

			assert.Equal(t,
				tt.want,
				got,
				"NewListRoleRepo() = %v, want = %v",
				got,
				tt.want)
		})
	}
}

func Test_listRoleRepo_ListRole(t *testing.T) {
	type fields struct {
		store ListRoleStore
	}
	type args struct {
		ctx context.Context
	}

	mockRole := new(mockListRoleStore)

	roles := []rolemodel.SimpleRole{
		{
			Id:   mock.Anything,
			Name: mock.Anything,
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []rolemodel.SimpleRole
		wantErr bool
	}{
		{
			name: "list role failed because can not get data from database",
			fields: fields{
				store: mockRole,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRole.
					On(
						"ListRole",
						context.Background()).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "list role successfully",
			fields: fields{
				store: mockRole,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockRole.
					On(
						"ListRole",
						context.Background()).
					Return(roles, nil).
					Once()
			},
			want:    roles,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listRoleRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := biz.ListRole(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListRole() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListRole() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListRole() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
