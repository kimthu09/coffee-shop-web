package productbiz

import (
	"coffee_shop_management_backend/common"
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
	repo CreateToppingRepo
}

func NewCreateToppingBiz(
	repo CreateToppingRepo) *createToppingBiz {
	return &createToppingBiz{
		repo: repo,
	}
}

func (biz *createToppingBiz) CreateTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckIngredient(ctx, data); err != nil {
		return err
	}

	if err := handleId(data); err != nil {
		return err
	}

	if err := biz.repo.StoreTopping(ctx, data); err != nil {
		return err
	}
	return nil
}

func handleId(data *productmodel.ToppingCreate) error {
	if err := handleToppingId(data); err != nil {
		return err
	}

	if err := handleRecipeId(data); err != nil {
		return err
	}

	return nil
}

func handleToppingId(data *productmodel.ToppingCreate) error {
	idTopping, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idTopping
	return nil
}

func handleRecipeId(data *productmodel.ToppingCreate) error {
	idRecipe, err := common.IdProcess(&data.RecipeId)
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
