package customerrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListCustomerStore struct {
	mock.Mock
}

func (m *mockListCustomerStore) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]customermodel.Customer, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	return args.Get(0).([]customermodel.Customer), args.Error(1)
}

func TestNewListCustomerRepo(t *testing.T) {
	type args struct {
		store ListCustomerStore
	}

	store := new(mockListCustomerStore)

	tests := []struct {
		name string
		args args
		want *listCustomerRepo
	}{
		{
			name: "Create object has type ListCustomerRepo",
			args: args{
				store: store,
			},
			want: &listCustomerRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListCustomerRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListCustomerRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listCustomerRepo_ListCustomer(t *testing.T) {
	type fields struct {
		store ListCustomerStore
	}
	type args struct {
		ctx    context.Context
		filter *customermodel.Filter
		paging *common.Paging
	}

	store := new(mockListCustomerStore)

	filterCustomer := &customermodel.Filter{
		SearchKey: "",
	}

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	listCustomers := make([]customermodel.Customer, 0)
	var emptyListCustomers []customermodel.Customer

	mockErr := assert.AnError

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []customermodel.Customer
		wantErr bool
	}{
		{
			name: "List customers failed because can not get data from database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterCustomer,
				paging: paging,
			},
			mock: func() {
				store.
					On("ListCustomer",
						context.Background(),
						filterCustomer,
						[]string{"id", "name", "email", "phone"},
						paging).
					Return(emptyListCustomers, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List customers successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterCustomer,
				paging: paging,
			},
			mock: func() {
				store.
					On("ListCustomer",
						context.Background(),
						filterCustomer,
						[]string{"id", "name", "email", "phone"},
						paging).
					Return(listCustomers, nil).
					Once()
			},
			want:    listCustomers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listCustomerRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListCustomer(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListCustomer() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListCustomer() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListCustomer() = %v, want %v", got, tt.want)
			}

		})
	}
}
