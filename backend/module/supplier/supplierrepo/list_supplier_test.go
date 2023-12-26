package supplierrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListSupplierStore struct {
	mock.Mock
}

func (m *mockListSupplierStore) ListSupplier(
	ctx context.Context,
	filter *filter.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	return args.Get(0).([]suppliermodel.Supplier), args.Error(1)
}

func TestNewListSupplierRepo(t *testing.T) {
	type args struct {
		store ListSupplierStore
	}

	store := new(mockListSupplierStore)

	tests := []struct {
		name string
		args args
		want *listSupplierRepo
	}{
		{
			name: "Create object has type ListSupplierRepo",
			args: args{
				store: store,
			},
			want: &listSupplierRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListSupplierRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListSupplierRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listSupplierRepo_ListSupplier(t *testing.T) {
	type fields struct {
		store ListSupplierStore
	}
	type args struct {
		ctx    context.Context
		filter *filter.Filter
		paging *common.Paging
	}

	store := new(mockListSupplierStore)

	filterSupplier := &filter.Filter{
		SearchKey: "",
		MinDebt:   nil,
		MaxDebt:   nil,
	}

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	listSuppliers := make([]suppliermodel.Supplier, 0)
	var emptyListSuppliers []suppliermodel.Supplier

	mockErr := assert.AnError

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []suppliermodel.Supplier
		wantErr bool
	}{
		{
			name: "List suppliers failed because can not get data from database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterSupplier,
				paging: paging,
			},
			mock: func() {
				store.
					On("ListSupplier",
						context.Background(),
						filterSupplier,
						[]string{"id", "name", "email", "phone"},
						paging).
					Return(emptyListSuppliers, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List suppliers successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterSupplier,
				paging: paging,
			},
			mock: func() {
				store.
					On("ListSupplier",
						context.Background(),
						filterSupplier,
						[]string{"id", "name", "email", "phone"},
						paging).
					Return(listSuppliers, nil).
					Once()
			},
			want:    listSuppliers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listSupplierRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListSupplier(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListSupplier() = %v, want %v", got, tt.want)
			}

		})
	}
}
