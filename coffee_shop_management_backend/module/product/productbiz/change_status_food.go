package productbiz

import (
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
	repo ChangeStatusFoodRepo
}

func NewChangeStatusFoodBiz(repo ChangeStatusFoodRepo) *changeStatusFoodBiz {
	return &changeStatusFoodBiz{repo: repo}
}

func (biz *changeStatusFoodBiz) ChangeStatusFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdate) error {

	if err := biz.repo.ChangeStatusFood(ctx, id, data); err != nil {
		return err
	}
	return nil
}
