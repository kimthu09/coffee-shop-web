package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type UpdateInfoSupplierStore interface {
	UpdateSupplierInfo(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateInfo,
	) error
}

type updateInfoSupplierRepo struct {
	store UpdateInfoSupplierStore
}

func NewUpdateInfoSupplierRepo(store UpdateInfoSupplierStore) *updateInfoSupplierRepo {
	return &updateInfoSupplierRepo{store: store}
}

func (repo *updateInfoSupplierRepo) UpdateSupplierInfo(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.SupplierUpdateInfo) error {
	if err := repo.store.UpdateSupplierInfo(ctx, supplierId, data); err != nil {
		return err
	}
	return nil
}
