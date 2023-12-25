package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ListFoodRepo interface {
	ListFood(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
	) ([]productmodel.Food, error)
}

type listFoodBiz struct {
	repo      ListFoodRepo
	requester middleware.Requester
}

func NewListFoodBiz(
	repo ListFoodRepo,
	requester middleware.Requester) *listFoodBiz {
	return &listFoodBiz{repo: repo, requester: requester}
}

func (biz *listFoodBiz) ListFood(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging) ([]productmodel.Food, error) {
	if !biz.requester.IsHasFeature(common.FoodViewFeatureCode) {
		return nil, productmodel.ErrFoodViewNoPermission
	}

	result, err := biz.repo.ListFood(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
