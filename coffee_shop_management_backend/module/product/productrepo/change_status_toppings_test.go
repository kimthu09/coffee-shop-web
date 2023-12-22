package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeStatusToppingsStore struct {
	mock.Mock
}

func (m *mockChangeStatusToppingsStore) UpdateStatusTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdateStatus) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewChangeStatusToppingsRepo(t *testing.T) {
	type args struct {
		store ChangeStatusToppingsStore
	}

	mockStore := new(mockChangeStatusToppingsStore)

	tests := []struct {
		name string
		args args
		want *changeStatusToppingsRepo
	}{
		{
			name: "Create object has type ChangeStatusToppingsRepo",
			args: args{
				store: mockStore,
			},
			want: &changeStatusToppingsRepo{
				store: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusToppingsRepo(
				tt.args.store,
			)

			assert.Equal(t, tt.want, got, "NewChangeStatusToppingsRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusToppingsRepo_ChangeStatusToppings(t *testing.T) {
	type fields struct {
		store ChangeStatusToppingsStore
	}
	type args struct {
		ctx  context.Context
		data []productmodel.ToppingUpdateStatus
	}

	mockStore := new(mockChangeStatusToppingsStore)

	active := true
	fakeToppingUpdateStatus := []productmodel.ToppingUpdateStatus{
		{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: "topping1",
				IsActive:  &active,
			},
		},
	}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Change status toppings failed due to store error",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeToppingUpdateStatus,
			},
			mock: func() {
				mockStore.
					On(
						"UpdateStatusTopping",
						context.Background(),
						"topping1", &fakeToppingUpdateStatus[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status toppings successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeToppingUpdateStatus,
			},
			mock: func() {
				mockStore.
					On(
						"UpdateStatusTopping",
						context.Background(),
						"topping1",
						&fakeToppingUpdateStatus[0]).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusToppingsRepo{
				store: tt.fields.store,
			}

			tt.mock()

			err := repo.ChangeStatusToppings(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "ChangeStatusToppings() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ChangeStatusToppings() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
