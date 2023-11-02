package productbiz

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusFoodStore interface {
	UpdateFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdate,
	) error
}

type changeStatusFoodBiz struct {
	store ChangeStatusFoodStore
}

func NewChangeStatusFoodBiz(store ChangeStatusFoodStore) *changeStatusFoodBiz {
	return &changeStatusFoodBiz{store: store}
}

func (biz *changeStatusFoodBiz) ChangeStatusFood(
	ctx context.Context,
	id string,
	toValue bool) error {

	var data productmodel.FoodUpdate
	data.IsActive = &toValue

	err := biz.store.UpdateFood(ctx, id, &data)

	return err
}
