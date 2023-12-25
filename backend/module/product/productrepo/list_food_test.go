package productrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListFoodStore struct {
	mock.Mock
}

func (m *mockListFoodStore) ListFood(
	ctx context.Context,
	filter *productmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
) ([]productmodel.Food, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	return args.Get(0).([]productmodel.Food), args.Error(1)
}

func TestNewListFoodRepo(t *testing.T) {
	type args struct {
		store ListFoodStore
	}
	store := new(mockListFoodStore)

	tests := []struct {
		name string
		args args
		want *listFoodRepo
	}{
		{
			name: "Create object has type ListFoodRepo",
			args: args{
				store: store,
			},
			want: &listFoodRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListFoodRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListFoodRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listFoodRepo_ListFood(t *testing.T) {
	type fields struct {
		store ListFoodStore
	}
	type args struct {
		ctx    context.Context
		filter *productmodel.Filter
		paging *common.Paging
	}

	mockStore := new(mockListFoodStore)

	filterFood := &productmodel.Filter{}

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	mockErr := assert.AnError

	listFoods := make([]productmodel.Food, 0)
	var emptyListFoods []productmodel.Food

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []productmodel.Food
		wantErr bool
	}{
		{
			name: "List foods failed because can not get data from the mockStore",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterFood,
				paging: paging,
			},
			mock: func() {
				mockStore.
					On("ListFood",
						context.Background(),
						filterFood,
						[]string{"id", "name"},
						paging).
					Return(emptyListFoods, mockErr).
					Once()
			},
			want:    listFoods,
			wantErr: true,
		},
		{
			name: "List foods successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterFood,
				paging: paging,
			},
			mock: func() {
				mockStore.
					On("ListFood",
						context.Background(),
						filterFood,
						[]string{"id", "name"},
						paging).
					Return(listFoods, nil).
					Once()
			},
			want:    listFoods,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listFoodRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListFood(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
