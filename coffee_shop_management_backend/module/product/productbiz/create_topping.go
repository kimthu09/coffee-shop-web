package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type CreateToppingRepo interface {
	CheckIngredient(
		ctx context.Context,
		data *productmodel.ToppingCreate,
	) error
	StoreTopping(
		ctx context.Context,
		data *productmodel.ToppingCreate,
	) error
}

type createToppingBiz struct {
	gen       generator.IdGenerator
	repo      CreateToppingRepo
	requester middleware.Requester
}

func NewCreateToppingBiz(
	gen generator.IdGenerator,
	repo CreateToppingRepo,
	requester middleware.Requester) *createToppingBiz {
	return &createToppingBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *createToppingBiz) CreateTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	if !biz.requester.IsHasFeature(common.ToppingCreateFeatureCode) {
		return productmodel.ErrToppingCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckIngredient(ctx, data); err != nil {
		return err
	}

	if err := handleId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.StoreTopping(ctx, data); err != nil {
		return err
	}
	return nil
}

func handleId(gen generator.IdGenerator, data *productmodel.ToppingCreate) error {
	if err := handleToppingId(gen, data); err != nil {
		return err
	}

	if err := handleRecipeId(gen, data); err != nil {
		return err
	}

	return nil
}

func handleToppingId(gen generator.IdGenerator, data *productmodel.ToppingCreate) error {
	idTopping, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idTopping
	return nil
}

func handleRecipeId(gen generator.IdGenerator, data *productmodel.ToppingCreate) error {
	idRecipe, err := gen.IdProcess(&data.RecipeId)
	if err != nil {
		return err
	}

	data.RecipeId = *idRecipe
	data.Recipe.Id = *idRecipe
	for i := range data.Recipe.Details {
		data.Recipe.Details[i].RecipeId = *idRecipe
	}
	return nil
}
