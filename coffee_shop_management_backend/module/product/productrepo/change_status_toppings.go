package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusToppingsStore interface {
	UpdateStatusTopping(
		ctx context.Context,
		id string,
		data *productmodel.ToppingUpdateStatus,
	) error
}

type changeStatusToppingsRepo struct {
	store ChangeStatusToppingsStore
}

func NewChangeStatusToppingsRepo(store ChangeStatusToppingsStore) *changeStatusToppingsRepo {
	return &changeStatusToppingsRepo{store: store}
}

func (biz *changeStatusToppingsRepo) ChangeStatusToppings(
	ctx context.Context,
	data []productmodel.ToppingUpdateStatus) error {
	for _, v := range data {
		if err := biz.store.UpdateStatusTopping(
			ctx, v.ProductId, &v); err != nil {
			return err
		}
	}

	return nil
}
