package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
)

type CreateToppingStore interface {
	CreateTopping(
		ctx context.Context,
		data *productmodel.ToppingCreate,
	) error
}

type CreateRecipeStore interface {
	CreateRecipe(
		ctx context.Context,
		data *recipemodel.RecipeCreate,
	) error
}

type CreateListRecipeDetailStore interface {
	CreateListRecipeDetail(
		ctx context.Context,
		data []recipedetailmodel.RecipeDetailCreate,
	) error
}

type createToppingRepo struct {
	toppingStore      CreateToppingStore
	recipeStore       CreateRecipeStore
	recipeDetailStore CreateListRecipeDetailStore
}

func NewCreateToppingRepo(
	toppingStore CreateToppingStore,
	recipeStore CreateRecipeStore,
	recipeDetailStore CreateListRecipeDetailStore) *createToppingRepo {
	return &createToppingRepo{
		toppingStore:      toppingStore,
		recipeStore:       recipeStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (repo *createToppingRepo) StoreTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	if err := repo.recipeStore.CreateRecipe(ctx, data.Recipe); err != nil {
		return err
	}

	if err := repo.recipeDetailStore.CreateListRecipeDetail(ctx, data.Recipe.Details); err != nil {
		return err
	}

	if err := repo.toppingStore.CreateTopping(ctx, data); err != nil {
		return err
	}

	return nil
}
