package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type SeeDetailToppingRepo interface {
	SeeDetailTopping(
		ctx context.Context,
		toppingId string,
	) (*productmodel.Topping, error)
}

type seeDetailToppingBiz struct {
	repo      SeeDetailToppingRepo
	requester middleware.Requester
}

func NewSeeDetailToppingBiz(
	repo SeeDetailToppingRepo,
	requester middleware.Requester) *seeDetailToppingBiz {
	return &seeDetailToppingBiz{repo: repo, requester: requester}
}

func (biz *seeDetailToppingBiz) SeeDetailTopping(
	ctx context.Context,
	toppingId string) (*productmodel.Topping, error) {
	if !biz.requester.IsHasFeature(common.ToppingViewFeatureCode) {
		return nil, productmodel.ErrToppingViewNoPermission
	}

	result, err := biz.repo.SeeDetailTopping(ctx, toppingId)

	if err != nil {
		return nil, err
	}

	return result, nil
}
