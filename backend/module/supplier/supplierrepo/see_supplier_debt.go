package supplierrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type ListSupplierDebtStore interface {
	ListSupplierDebt(
		ctx context.Context,
		supplierId string,
		filterSupplierDebt *filter.SupplierDebtFilter,
		paging *common.Paging,
		moreKeys ...string) ([]supplierdebtmodel.SupplierDebt, error)
}

type seeSupplierDebtRepo struct {
	debtStore ListSupplierDebtStore
}

func NewSeeSupplierDebtRepo(
	debtStore ListSupplierDebtStore) *seeSupplierDebtRepo {
	return &seeSupplierDebtRepo{
		debtStore: debtStore,
	}
}

func (biz *seeSupplierDebtRepo) SeeSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	debts, errDebts := biz.debtStore.ListSupplierDebt(
		ctx,
		supplierId,
		filterSupplierDebt,
		paging,
		"CreatedByUser",
	)
	if errDebts != nil {
		return nil, errDebts
	}

	return debts, nil
}
