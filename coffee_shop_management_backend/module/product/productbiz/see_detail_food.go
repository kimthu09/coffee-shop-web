package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type SeeDetailFoodRepo interface {
	SeeDetailFood(
		ctx context.Context,
		foodId string) (*productmodel.Food, error)
}

type seeDetailFoodBiz struct {
	repo      SeeDetailFoodRepo
	requester middleware.Requester
}

func NewSeeDetailFoodBiz(
	repo SeeDetailFoodRepo,
	requester middleware.Requester) *seeDetailFoodBiz {
	return &seeDetailFoodBiz{repo: repo, requester: requester}
}

func (biz *seeDetailFoodBiz) SeeDetailFood(
	ctx context.Context,
	foodId string) (*productmodel.Food, error) {
	if !biz.requester.IsHasFeature(common.FoodViewFeatureCode) {
		return nil, productmodel.ErrFoodViewNoPermission
	}

	result, err := biz.repo.SeeDetailFood(ctx, foodId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
