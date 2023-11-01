package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type UpdateToppingStore interface {
	FindTopping(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*productmodel.Topping, error)
	UpdateTopping(
		ctx context.Context,
		id string,
		data *productmodel.ToppingUpdate) error
}

type updateToppingBiz struct {
	store UpdateToppingStore
}

func NewUpdateToppingBiz(store UpdateToppingStore) *updateToppingBiz {
	return &updateToppingBiz{store: store}
}

func (biz *updateToppingBiz) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdate) error {

	if err := data.Validate(); err != nil {
		return err
	}

	result, err := biz.store.FindTopping(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if !result.IsActive {
		return common.ErrNoPermission(productmodel.ErrProductInactive)
	}

	errStorage := biz.store.UpdateTopping(ctx, id, data)

	return errStorage
}
