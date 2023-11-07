package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusFoodRepo interface {
	ChangeStatusFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdate,
	) error
}

type changeStatusFoodBiz struct {
	repo      ChangeStatusFoodRepo
	requester middleware.Requester
}

func NewChangeStatusFoodBiz(
	repo ChangeStatusFoodRepo,
	requester middleware.Requester) *changeStatusFoodBiz {
	return &changeStatusFoodBiz{repo: repo, requester: requester}
}

func (biz *changeStatusFoodBiz) ChangeStatusFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdate) error {
	if !biz.requester.IsHasFeature(common.FoodUpdateStatusFeatureCode) {
		return productmodel.ErrFoodChangeStatusNoPermission
	}

	if err := biz.repo.ChangeStatusFood(ctx, id, data); err != nil {
		return err
	}
	return nil
}
