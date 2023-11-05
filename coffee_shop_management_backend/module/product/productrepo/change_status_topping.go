package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type ChangeStatusToppingStore interface {
	UpdateTopping(
		ctx context.Context,
		id string,
		data *productmodel.ToppingUpdate,
	) error
}

type changeStatusToppingRepo struct {
	store ChangeStatusToppingStore
}

func NewChangeStatusToppingRepo(store ChangeStatusToppingStore) *changeStatusToppingRepo {
	return &changeStatusToppingRepo{store: store}
}

func (biz *changeStatusToppingRepo) ChangeStatusTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdate) error {

	if err := biz.store.UpdateTopping(ctx, id, data); err != nil {
		return err
	}

	return nil
}
