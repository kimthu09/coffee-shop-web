package ingredientbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

type CreateIngredientStore interface {
	CreateIngredient(
		ctx context.Context,
		data *ingredientmodel.IngredientCreate,
	) error
}

type createIngredientBiz struct {
	store CreateIngredientStore
}

func NewCreateIngredientBiz(store CreateIngredientStore) *createIngredientBiz {
	return &createIngredientBiz{store: store}
}

func (biz *createIngredientBiz) CreateIngredient(
	ctx context.Context,
	data *ingredientmodel.IngredientCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}

	idAddress, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress

	if err := biz.store.CreateIngredient(ctx, data); err != nil {
		return err
	}

	return nil
}
