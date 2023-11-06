package supplierrepo

import (
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type PaySupplierStore interface {
	GetDebtSupplier(
		ctx context.Context,
		supplierId string) (*float32, error)
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
	requester         middleware.Requester
}

func NewUpdatePayRepo(
	supplierStore PaySupplierStore,
	supplierDebtStore CreateSupplierDebtStore) *paySupplierRepo {
	return &paySupplierRepo{
		supplierStore:     supplierStore,
		supplierDebtStore: supplierDebtStore,
	}
}

func (repo *paySupplierRepo) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*float32, error) {
	debtCurrent, err := repo.supplierStore.GetDebtSupplier(
		ctx,
		supplierId)
	if err != nil {
		return nil, err
	}

	return debtCurrent, nil
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
