package supplierdebtbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type ListSupplierDebtStore interface {
	ListSupplierDebt(
		ctx context.Context,
		supplierId string,
		paging *common.Paging,
	) ([]supplierdebtmodel.SupplierDebt, error)
}

type listSupplierDebtBiz struct {
	store     ListSupplierDebtStore
	requester middleware.Requester
}

func NewListSupplierDebtBiz(
	store ListSupplierDebtStore,
	requester middleware.Requester) *listSupplierDebtBiz {
	return &listSupplierDebtBiz{store: store, requester: requester}
}

func (biz *listSupplierDebtBiz) ListSupplierDebt(
	ctx context.Context,
	supplierId string,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, supplierdebtmodel.ErrSupplierDebtViewNoPermission
	}

	result, err := biz.store.ListSupplierDebt(ctx, supplierId, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
