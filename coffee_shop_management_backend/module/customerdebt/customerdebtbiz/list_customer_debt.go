package customerdebtbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"context"
)

type ListCustomerDebtStore interface {
	ListCustomerDebt(
		ctx context.Context,
		customerId string,
		paging *common.Paging,
	) ([]customerdebtmodel.CustomerDebt, error)
}

type listCustomerDebtBiz struct {
	store     ListCustomerDebtStore
	requester middleware.Requester
}

func NewListCustomerDebtBiz(
	store ListCustomerDebtStore,
	requester middleware.Requester) *listCustomerDebtBiz {
	return &listCustomerDebtBiz{store: store, requester: requester}
}

func (biz *listCustomerDebtBiz) ListCustomerDebt(
	ctx context.Context,
	customerId string,
	paging *common.Paging) ([]customerdebtmodel.CustomerDebt, error) {
	if !biz.requester.IsHasFeature(common.CustomerViewFeatureCode) {
		return nil, customerdebtmodel.ErrCustomerDebtViewNoPermission
	}

	result, err := biz.store.ListCustomerDebt(ctx, customerId, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
