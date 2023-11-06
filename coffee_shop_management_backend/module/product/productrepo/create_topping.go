package productrepo

import (
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
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

type CheckIngredientStore interface {
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientmodel.Ingredient, error)
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
	ingredientStore   CheckIngredientStore
	recipeDetailStore CreateListRecipeDetailStore
}

func NewCreateToppingRepo(
	toppingStore CreateToppingStore,
	recipeStore CreateRecipeStore,
	ingredientStore CheckIngredientStore,
	recipeDetailStore CreateListRecipeDetailStore) *createToppingRepo {
	return &createToppingRepo{
		toppingStore:      toppingStore,
		recipeStore:       recipeStore,
		ingredientStore:   ingredientStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (repo *createToppingRepo) CheckIngredient(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	for _, recipeDetail := range data.Recipe.Details {
		if _, err := repo.ingredientStore.FindIngredient(
			ctx,
			map[string]interface{}{"id": recipeDetail.IngredientId},
		); err != nil {
			return err
		}
	}
	return nil
}

func (repo *createToppingRepo) StoreTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	if err := repo.toppingStore.CreateTopping(ctx, data); err != nil {
		return err
	}

	if err := repo.recipeStore.CreateRecipe(ctx, data.Recipe); err != nil {
		return err
	}

	if err := repo.recipeDetailStore.CreateListRecipeDetail(ctx, data.Recipe.Details); err != nil {
		return err
	}

	return nil
}
