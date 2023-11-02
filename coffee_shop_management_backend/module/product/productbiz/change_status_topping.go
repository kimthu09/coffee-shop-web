package productbiz

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

type changeStatusToppingBiz struct {
	store ChangeStatusToppingStore
}

func NewChangeStatusToppingBiz(store ChangeStatusToppingStore) *changeStatusToppingBiz {
	return &changeStatusToppingBiz{store: store}
}

func (biz *changeStatusToppingBiz) ChangeStatusTopping(
	ctx context.Context,
	id string,
	toValue bool) error {

	var data productmodel.ToppingUpdate
	data.IsActive = &toValue

	err := biz.store.UpdateTopping(ctx, id, &data)

	return err
}
