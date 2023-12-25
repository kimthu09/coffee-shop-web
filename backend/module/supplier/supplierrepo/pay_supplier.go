package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type PaySupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*suppliermodel.Supplier, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate) error
}

type paySupplierRepo struct {
	supplierStore     PaySupplierStore
	supplierDebtStore CreateSupplierDebtStore
}

func NewPaySupplierRepo(
	supplierStore PaySupplierStore,
	supplierDebtStore CreateSupplierDebtStore) *paySupplierRepo {
	return &paySupplierRepo{
		supplierStore:     supplierStore,
		supplierDebtStore: supplierDebtStore,
	}
}

func (repo *paySupplierRepo) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*int, error) {
	supplier, err := repo.supplierStore.FindSupplier(
		ctx, map[string]interface{}{"id": supplierId},
	)
	if err != nil {
		return nil, err
	}

	return &supplier.Debt, nil
}

func (repo *paySupplierRepo) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
	if err := repo.supplierDebtStore.CreateSupplierDebt(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *paySupplierRepo) UpdateDebtSupplier(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.SupplierUpdateDebt) error {
	if err := repo.supplierStore.UpdateSupplierDebt(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
