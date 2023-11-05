package productbiz

import (
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
	repo ChangeStatusToppingRepo
}

func NewChangeStatusToppingBiz(repo ChangeStatusToppingRepo) *changeStatusToppingBiz {
	return &changeStatusToppingBiz{repo: repo}
}

func (biz *changeStatusToppingBiz) ChangeStatusTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdate) error {
	if err := biz.repo.ChangeStatusTopping(ctx, id, data); err != nil {
		return err
	}
	return nil
}
