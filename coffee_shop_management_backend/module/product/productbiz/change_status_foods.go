package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusFoodsRepo interface {
	ChangeStatusFoods(
		ctx context.Context,
		data []productmodel.FoodUpdateStatus) error
}

type changeStatusFoodsBiz struct {
	repo      ChangeStatusFoodsRepo
	requester middleware.Requester
}

func NewChangeStatusFoodsBiz(
	repo ChangeStatusFoodsRepo,
	requester middleware.Requester) *changeStatusFoodsBiz {
	return &changeStatusFoodsBiz{repo: repo, requester: requester}
}

func (biz *changeStatusFoodsBiz) ChangeStatusFoods(
	ctx context.Context,
	data []productmodel.FoodUpdateStatus) error {
	if !biz.requester.IsHasFeature(common.FoodUpdateStatusFeatureCode) {
		return productmodel.ErrFoodChangeStatusNoPermission
	}

	for _, v := range data {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	if err := biz.repo.ChangeStatusFoods(ctx, data); err != nil {
		return err
	}
	return nil
}
