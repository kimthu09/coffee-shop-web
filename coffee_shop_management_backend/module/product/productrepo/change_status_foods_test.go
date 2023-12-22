package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockChangeStatusFoodsStore struct {
	mock.Mock
}

func (m *mockChangeStatusFoodsStore) UpdateStatusFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdateStatus) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewChangeStatusFoodsRepo(t *testing.T) {
	type args struct {
		store ChangeStatusFoodsStore
	}

	mockStore := new(mockChangeStatusFoodsStore)

	tests := []struct {
		name string
		args args
		want *changeStatusFoodsRepo
	}{
		{
			name: "Create object has type ChangeStatusFoodsRepo",
			args: args{
				store: mockStore,
			},
			want: &changeStatusFoodsRepo{
				store: mockStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewChangeStatusFoodsRepo(
				tt.args.store,
			)

			assert.Equal(t, tt.want, got, "NewChangeStatusFoodsRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_changeStatusFoodsRepo_ChangeStatusFoods(t *testing.T) {
	type fields struct {
		store ChangeStatusFoodsStore
	}
	type args struct {
		ctx  context.Context
		data []productmodel.FoodUpdateStatus
	}

	mockStore := new(mockChangeStatusFoodsStore)

	active := true
	fakeFoodUpdateStatus := []productmodel.FoodUpdateStatus{
		{
			ProductUpdateStatus: &productmodel.ProductUpdateStatus{
				ProductId: "food1",
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
			name: "Change status foods failed due to store error",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodUpdateStatus,
			},
			mock: func() {
				mockStore.
					On(
						"UpdateStatusFood",
						context.Background(),
						"food1", &fakeFoodUpdateStatus[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Change status foods successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodUpdateStatus,
			},
			mock: func() {
				mockStore.
					On(
						"UpdateStatusFood",
						context.Background(),
						"food1",
						&fakeFoodUpdateStatus[0]).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &changeStatusFoodsRepo{
				store: tt.fields.store,
			}

			tt.mock()

			err := repo.ChangeStatusFoods(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "ChangeStatusFoods() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ChangeStatusFoods() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
