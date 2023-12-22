package productrepo

import (
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

type CreateFoodStore interface {
	CreateFood(
		ctx context.Context,
		data *productmodel.FoodCreate,
	) error
}

type CreateCategoryFoodStore interface {
	CreateCategoryFood(
		ctx context.Context,
		data *categoryfoodmodel.CategoryFoodCreate,
	) error
}

type UpdateCategoryStore interface {
	UpdateAmountProductCategory(
		ctx context.Context,
		id string,
		data *categorymodel.CategoryUpdateAmountProduct,
	) error
}

type CreateSizeFoodStore interface {
	CreateSizeFood(
		ctx context.Context,
		data *sizefoodmodel.SizeFoodCreate,
	) error
}

type createFoodRepo struct {
	foodStore         CreateFoodStore
	categoryFoodStore CreateCategoryFoodStore
	categoryStore     UpdateCategoryStore
	sizeFoodStore     CreateSizeFoodStore
	recipeStore       CreateRecipeStore
	recipeDetailStore CreateListRecipeDetailStore
}

func NewCreateFoodRepo(
	foodStore CreateFoodStore,
	categoryFoodStore CreateCategoryFoodStore,
	categoryStore UpdateCategoryStore,
	sizeFoodStore CreateSizeFoodStore,
	recipeStore CreateRecipeStore,
	recipeDetailStore CreateListRecipeDetailStore) *createFoodRepo {
	return &createFoodRepo{
		foodStore:         foodStore,
		categoryFoodStore: categoryFoodStore,
		categoryStore:     categoryStore,
		sizeFoodStore:     sizeFoodStore,
		recipeStore:       recipeStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (repo *createFoodRepo) CreateFood(
	ctx context.Context,
	data *productmodel.FoodCreate) error {
	if err := repo.foodStore.CreateFood(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *createFoodRepo) HandleCategoryFood(
	ctx context.Context,
	foodId string,
	data *productmodel.FoodCreate) error {
	for _, categoryId := range data.Categories {
		if err := repo.createCategoryFood(ctx, foodId, categoryId); err != nil {
			return err
		}
		if err := repo.updateAmountProductCategory(ctx, categoryId); err != nil {
			return err
		}
	}
	return nil
}

func (repo *createFoodRepo) createCategoryFood(
	ctx context.Context,
	foodId string,
	categoryId string) error {
	categoryFoodCreate := categoryfoodmodel.CategoryFoodCreate{
		FoodId:     foodId,
		CategoryId: categoryId,
	}
	if err := repo.categoryFoodStore.CreateCategoryFood(
		ctx,
		&categoryFoodCreate); err != nil {
		return err
	}
	return nil
}

func (repo *createFoodRepo) updateAmountProductCategory(
	ctx context.Context,
	categoryId string) error {
	amount := 1
	categoryUpdateAmountProduct := categorymodel.CategoryUpdateAmountProduct{
		AmountProduct: &amount,
	}
	if err := repo.categoryStore.UpdateAmountProductCategory(
		ctx,
		categoryId,
		&categoryUpdateAmountProduct); err != nil {
		return err
	}
	return nil
}

func (repo *createFoodRepo) HandleSizeFood(
	ctx context.Context,
	data *productmodel.FoodCreate) error {
	for _, value := range data.Sizes {
		if err := repo.createRecipe(ctx, *value.Recipe); err != nil {
			return err
		}

		if err := repo.createSizeFood(ctx, value); err != nil {
			return err
		}

		if err := repo.createRecipeDetails(ctx, value.Recipe.Details); err != nil {
			return err
		}
	}
	return nil
}

func (repo *createFoodRepo) createSizeFood(
	ctx context.Context,
	sizeFood sizefoodmodel.SizeFoodCreate) error {
	if err := repo.sizeFoodStore.CreateSizeFood(
		ctx,
		&sizeFood); err != nil {
		return err
	}
	return nil
}

func (repo *createFoodRepo) createRecipe(
	ctx context.Context,
	recipeCreate recipemodel.RecipeCreate) error {
	if err := repo.recipeStore.CreateRecipe(
		ctx,
		&recipeCreate); err != nil {
		return err
	}
	return nil
}

func (repo *createFoodRepo) createRecipeDetails(
	ctx context.Context,
	details []recipedetailmodel.RecipeDetailCreate) error {
	if err := repo.recipeDetailStore.CreateListRecipeDetail(ctx, details); err != nil {
		return err
	}
	return nil
}
