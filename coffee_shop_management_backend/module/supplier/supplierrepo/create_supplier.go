package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type CreateSupplierStore interface {
	CreateSupplier(ctx context.Context, data *suppliermodel.SupplierCreate) error
}

type createSupplierRepo struct {
	store CreateSupplierStore
}

func NewCreateSupplierRepo(store CreateSupplierStore) *createSupplierRepo {
	return &createSupplierRepo{store: store}
}

func (biz *createSupplierRepo) CreateSupplier(
	ctx context.Context,
	data *suppliermodel.SupplierCreate) error {
	if err := biz.store.CreateSupplier(ctx, data); err != nil {
		return err
	}

	return nil
}
