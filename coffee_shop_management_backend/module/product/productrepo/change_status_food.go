package productrepo

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

type changeStatusFoodRepo struct {
	store ChangeStatusFoodStore
}

func NewChangeStatusFoodRepo(store ChangeStatusFoodStore) *changeStatusFoodRepo {
	return &changeStatusFoodRepo{store: store}
}

func (biz *changeStatusFoodRepo) ChangeStatusFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdate) error {
	if err := biz.store.UpdateFood(ctx, id, data); err != nil {
		return err
	}
	return nil
}
