package supplierrepo

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type GetAllSupplierStore interface {
	GetAllSupplier(
		ctx context.Context,
		moreKeys ...string) ([]suppliermodel.SimpleSupplier, error)
}

type getAllSupplierRepo struct {
	store GetAllSupplierStore
}

func NewGetAllSupplierRepo(store GetAllSupplierStore) *getAllSupplierRepo {
	return &getAllSupplierRepo{store: store}
}

func (repo *getAllSupplierRepo) GetAllSupplier(
	ctx context.Context) ([]suppliermodel.SimpleSupplier, error) {
	result, err := repo.store.GetAllSupplier(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
