package supplierrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type ListSupplierStore interface {
	ListSupplier(
		ctx context.Context,
		filter *suppliermodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]suppliermodel.Supplier, error)
}

type listSupplierRepo struct {
	store ListSupplierStore
}

func NewListSupplierRepo(store ListSupplierStore) *listSupplierRepo {
	return &listSupplierRepo{store: store}
}

func (repo *listSupplierRepo) ListSupplier(
	ctx context.Context,
	filter *suppliermodel.Filter,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	result, err := repo.store.ListSupplier(
		ctx,
		filter,
		[]string{"id", "name", "email", "phone"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
