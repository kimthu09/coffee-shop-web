package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockPaySupplierStore struct {
	mock.Mock
}

func (m *mockPaySupplierStore) FindSupplier(
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

func (m *mockPaySupplierStore) UpdateSupplierDebt(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateDebt,
) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockCreateSupplierDebtStore struct {
	mock.Mock
}

func (m *mockCreateSupplierDebtStore) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate,
) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewUpdatePayRepo(t *testing.T) {
	type args struct {
		supplierStore     PaySupplierStore
		supplierDebtStore CreateSupplierDebtStore
	}

	supplierStore := new(mockPaySupplierStore)
	debtStore := new(mockCreateSupplierDebtStore)

	tests := []struct {
		name string
		args args
		want *paySupplierRepo
	}{
		{
			name: "Create object has type NewPaySupplierRepo",
			args: args{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			want: &paySupplierRepo{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPaySupplierRepo(tt.args.supplierStore, tt.args.supplierDebtStore)
			assert.Equal(t, tt.want, got, "NewPaySupplierRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_paySupplierRepo_CreateSupplierDebt(t *testing.T) {
	type fields struct {
		supplierStore     PaySupplierStore
		supplierDebtStore CreateSupplierDebtStore
	}
	type args struct {
		ctx  context.Context
		data *supplierdebtmodel.SupplierDebtCreate
	}

	supplierStore := new(mockPaySupplierStore)
	debtStore := new(mockCreateSupplierDebtStore)

	supplierDebtCreate := new(supplierdebtmodel.SupplierDebtCreate)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create supplier debt failed because can not save to database",
			fields: fields{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			args: args{
				ctx:  context.Background(),
				data: supplierDebtCreate,
			},
			mock: func() {
				debtStore.
					On("CreateSupplierDebt",
						context.Background(),
						supplierDebtCreate,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create supplier debt successfully",
			fields: fields{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			args: args{
				ctx:  context.Background(),
				data: supplierDebtCreate,
			},
			mock: func() {
				debtStore.
					On("CreateSupplierDebt",
						context.Background(),
						supplierDebtCreate,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &paySupplierRepo{
				supplierStore:     tt.fields.supplierStore,
				supplierDebtStore: tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.CreateSupplierDebt(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateSupplierDebt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_paySupplierRepo_GetDebtSupplier(t *testing.T) {
	type fields struct {
		supplierStore     PaySupplierStore
		supplierDebtStore CreateSupplierDebtStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
	}

	supplierStore := new(mockPaySupplierStore)
	debtStore := new(mockCreateSupplierDebtStore)

	supplierId := mock.Anything
	debt := -100000
	supplier := suppliermodel.Supplier{
		Id:   supplierId,
		Debt: debt,
	}

	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *int
		wantErr bool
	}{
		{
			name: "Get supplier debt failed",
			fields: fields{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				supplierStore.
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
			name: "Get supplier debt successfully",
			fields: fields{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: supplierId,
			},
			mock: func() {
				supplierStore.
					On("FindSupplier",
						context.Background(),
						map[string]interface{}{"id": supplierId},
						moreKeys,
					).
					Return(&supplier, nil).
					Once()
			},
			want:    &debt,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &paySupplierRepo{
				supplierStore:     tt.fields.supplierStore,
				supplierDebtStore: tt.fields.supplierDebtStore,
			}

			tt.mock()

			got, err := repo.GetDebtSupplier(tt.args.ctx, tt.args.supplierId)

			if tt.wantErr {
				assert.NotNil(t, err, "GetDebtSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetDebtSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetDebtSupplier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paySupplierRepo_UpdateDebtSupplier(t *testing.T) {
	type fields struct {
		supplierStore     PaySupplierStore
		supplierDebtStore CreateSupplierDebtStore
	}
	type args struct {
		ctx        context.Context
		supplierId string
		updateData *suppliermodel.SupplierUpdateDebt
	}

	supplierStore := new(mockPaySupplierStore)
	debtStore := new(mockCreateSupplierDebtStore)

	mockErr := errors.New(mock.Anything)

	debt := 100000
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount:    &debt,
		CreatedBy: mock.Anything,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update supplier debt failed because can not save to database",
			fields: fields{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: "supplier123",
				updateData: &supplierUpdateDebt,
			},
			mock: func() {
				supplierStore.
					On("UpdateSupplierDebt",
						context.Background(),
						"supplier123",
						&supplierUpdateDebt,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update supplier debt successfully",
			fields: fields{
				supplierStore:     supplierStore,
				supplierDebtStore: debtStore,
			},
			args: args{
				ctx:        context.Background(),
				supplierId: "supplier123",
				updateData: &supplierUpdateDebt,
			},
			mock: func() {
				supplierStore.
					On("UpdateSupplierDebt",
						context.Background(),
						"supplier123",
						&supplierUpdateDebt,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &paySupplierRepo{
				supplierStore:     tt.fields.supplierStore,
				supplierDebtStore: tt.fields.supplierDebtStore,
			}

			tt.mock()

			err := repo.UpdateDebtSupplier(tt.args.ctx, tt.args.supplierId, tt.args.updateData)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateDebtSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateDebtSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
