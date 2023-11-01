package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type CreateToppingStore interface {
	CreateTopping(
		ctx context.Context,
		data *productmodel.ToppingCreate,
	) error
}

type createToppingBiz struct {
	store CreateToppingStore
}

func NewCreateToppingBiz(store CreateToppingStore) *createToppingBiz {
	return &createToppingBiz{store: store}
}

func (biz *createToppingBiz) CreateTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}

	idAddress, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress

	errStorage := biz.store.CreateTopping(ctx, data)

	return errStorage
}
