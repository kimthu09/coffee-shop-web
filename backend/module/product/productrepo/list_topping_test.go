package productrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListToppingStore struct {
	mock.Mock
}

func (m *mockListToppingStore) ListTopping(
	ctx context.Context,
	filter *productmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
) ([]productmodel.Topping, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]productmodel.Topping), args.Error(1)
}

func TestNewListToppingRepo(t *testing.T) {
	type args struct {
		store ListToppingStore
	}
	store := new(mockListToppingStore)

	tests := []struct {
		name string
		args args
		want *listToppingRepo
	}{
		{
			name: "Create object has type ListToppingRepo",
			args: args{
				store: store,
			},
			want: &listToppingRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListToppingRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListToppingRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listToppingRepo_ListTopping(t *testing.T) {
	type fields struct {
		store ListToppingStore
	}
	type args struct {
		ctx    context.Context
		filter *productmodel.Filter
		paging *common.Paging
	}

	mockStore := new(mockListToppingStore)

	filterTopping := &productmodel.Filter{}

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	mockErr := assert.AnError

	listToppings := make([]productmodel.Topping, 0)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []productmodel.Topping
		wantErr bool
	}{
		{
			name: "List toppings failed because can not get data from the mockStore",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterTopping,
				paging: paging,
			},
			mock: func() {
				mockStore.
					On("ListTopping",
						context.Background(),
						filterTopping,
						[]string{"id", "name"},
						paging).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List toppings successfully",
			fields: fields{
				store: mockStore,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterTopping,
				paging: paging,
			},
			mock: func() {
				mockStore.
					On("ListTopping",
						context.Background(),
						filterTopping,
						[]string{"id", "name"},
						paging).
					Return(listToppings, nil).
					Once()
			},
			want:    listToppings,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listToppingRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListTopping(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListTopping() = %v, want %v", got, tt.want)
			}
		})
	}
}
