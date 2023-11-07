package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusToppingRepo interface {
	ChangeStatusTopping(
		ctx context.Context,
		id string,
		data *productmodel.ToppingUpdate,
	) error
}

type changeStatusToppingBiz struct {
	repo      ChangeStatusToppingRepo
	requester middleware.Requester
}

func NewChangeStatusToppingBiz(
	repo ChangeStatusToppingRepo,
	requester middleware.Requester) *changeStatusToppingBiz {
	return &changeStatusToppingBiz{repo: repo, requester: requester}
}

func (biz *changeStatusToppingBiz) ChangeStatusTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdate) error {
	if !biz.requester.IsHasFeature(common.ToppingUpdateStatusFeatureCode) {
		return productmodel.ErrToppingChangeStatusNoPermission
	}

	if err := biz.repo.ChangeStatusTopping(ctx, id, data); err != nil {
		return err
	}
	return nil
}
