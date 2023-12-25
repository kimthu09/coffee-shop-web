package ingredientbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
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
	gen       generator.IdGenerator
	store     CreateIngredientStore
	requester middleware.Requester
}

func NewCreateIngredientBiz(
	gen generator.IdGenerator,
	store CreateIngredientStore,
	requester middleware.Requester) *createIngredientBiz {
	return &createIngredientBiz{
		gen:       gen,
		store:     store,
		requester: requester,
	}
}

func (biz *createIngredientBiz) CreateIngredient(
	ctx context.Context,
	data *ingredientmodel.IngredientCreate) error {
	if !biz.requester.IsHasFeature(common.IngredientCreateFeatureCode) {
		return ingredientmodel.ErrIngredientCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	data.Round()

	idAddress, err := biz.gen.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress

	if err := biz.store.CreateIngredient(ctx, data); err != nil {
		return err
	}

	return nil
}
