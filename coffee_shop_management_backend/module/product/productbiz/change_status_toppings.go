package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusToppingsRepo interface {
	ChangeStatusToppings(
		ctx context.Context,
		data []productmodel.ToppingUpdateStatus) error
}

type changeStatusToppingsBiz struct {
	repo      ChangeStatusToppingsRepo
	requester middleware.Requester
}

func NewChangeStatusToppingsBiz(
	repo ChangeStatusToppingsRepo,
	requester middleware.Requester) *changeStatusToppingsBiz {
	return &changeStatusToppingsBiz{repo: repo, requester: requester}
}

func (biz *changeStatusToppingsBiz) ChangeStatusToppings(
	ctx context.Context,
	data []productmodel.ToppingUpdateStatus) error {
	if !biz.requester.IsHasFeature(common.ToppingUpdateStatusFeatureCode) {
		return productmodel.ErrToppingChangeStatusNoPermission
	}

	for _, v := range data {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	if err := biz.repo.ChangeStatusToppings(ctx, data); err != nil {
		return err
	}
	return nil
}
