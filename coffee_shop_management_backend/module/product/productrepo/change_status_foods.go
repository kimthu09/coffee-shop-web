package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusFoodsStore interface {
	UpdateStatusFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdateStatus,
	) error
}

type changeStatusFoodsRepo struct {
	store ChangeStatusFoodsStore
}

func NewChangeStatusFoodsRepo(store ChangeStatusFoodsStore) *changeStatusFoodsRepo {
	return &changeStatusFoodsRepo{store: store}
}

func (biz *changeStatusFoodsRepo) ChangeStatusFoods(
	ctx context.Context,
	data []productmodel.FoodUpdateStatus) error {
	for _, v := range data {
		if err := biz.store.UpdateStatusFood(
			ctx, v.ProductId, &v); err != nil {
			return err
		}
	}

	return nil
}
