package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ListToppingRepo interface {
	ListTopping(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
	) ([]productmodel.Topping, error)
}

type listToppingBiz struct {
	repo      ListToppingRepo
	requester middleware.Requester
}

func NewListToppingBiz(
	repo ListToppingRepo,
	requester middleware.Requester) *listToppingBiz {
	return &listToppingBiz{repo: repo, requester: requester}
}

func (biz *listToppingBiz) ListTopping(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging) ([]productmodel.Topping, error) {
	if !biz.requester.IsHasFeature(common.ToppingViewFeatureCode) {
		return nil, productmodel.ErrToppingViewNoPermission
	}

	result, err := biz.repo.ListTopping(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
