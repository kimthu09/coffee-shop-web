package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockFindSupplierStore struct {
	mock.Mock
}

func (m *mockFindSupplierStore) FindSupplier(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*suppliermodel.Supplier, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*suppliermodel.Supplier), args.Error(1)
}

func TestNewSeeSupplierDetailRepo(t *testing.T) {
	type args struct {
		supplierStore FindSupplierStore
	}

	store := new(mockFindSupplierStore)

	tests := []struct {
		name string
		args args
		want *seeSupplierDetailRepo
	}{
		{
			name: "Create object has type NewSeeSupplierDetailRepo",
			args: args{
				supplierStore: store,
			},
			want: &seeSupplierDetailRepo{
				supplierStore: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeSupplierDetailRepo(tt.args.supplierStore)
			assert.Equal(t, tt.want, got, "NewSeeSupplierDetailRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeSupplierDetailRepo_SeeSupplierDetail(t *testing.T) {
	type fields struct {
		supplierStore FindSupplierStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
	}

	store := new(mockFindSupplierStore)

	supplierId := mock.Anything
	mockSupplier := &suppliermodel.Supplier{
		Id:    supplierId,
		Name:  mock.Anything,
		Email: mock.Anything,
		Phone: mock.Anything,
		Debt:  -100000,
	}

	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *suppliermodel.Supplier
		wantErr bool
	}{
		{
			name: "See supplier detail failed because can not get data from database",
			fields: fields{
				supplierStore: store,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				store.
					On("FindSupplier",
						context.Background(),
						map[string]interface{}{"id": supplierId},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See supplier detail successfully",
			fields: fields{
				supplierStore: store,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				store.
					On("FindSupplier",
						context.Background(),
						map[string]interface{}{"id": supplierId},
						moreKeys,
					).
					Return(mockSupplier, nil).
					Once()
			},
			want:    mockSupplier,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &seeSupplierDetailRepo{
				supplierStore: tt.fields.supplierStore,
			}

			tt.mock()

			got, err := repo.SeeSupplierDetail(tt.args.ctx, tt.args.supplierId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeSupplierDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeSupplierDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeSupplierDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
