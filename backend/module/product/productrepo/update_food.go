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

type UpdateFoodStore interface {
	FindFood(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Food, error)
	UpdateFood(
		ctx context.Context,
		id string,
		data *productmodel.FoodUpdateInfo,
	) error
}

type CreateOrDeleteCategoryFoodStore interface {
	FindListCategories(
		ctx context.Context,
		foodId string,
	) ([]categorymodel.SimpleCategoryWithId, error)
	CreateCategoryFood(
		ctx context.Context,
		data *categoryfoodmodel.CategoryFoodCreate,
	) error
	DeleteCategoryFood(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type UpdateSizeFoodStore interface {
	FindListSizeFood(
		ctx context.Context,
		foodId string,
	) ([]sizefoodmodel.SizeFood, error)
	CreateSizeFood(
		ctx context.Context,
		data *sizefoodmodel.SizeFoodCreate,
	) error
	DeleteSizeFood(
		ctx context.Context,
		conditions map[string]interface{},
	) error
	UpdateSizeFood(
		ctx context.Context,
		foodId string,
		sizeId string,
		data *sizefoodmodel.SizeFoodUpdate,
	) error
}

type UpdateRecipeStore interface {
	CreateRecipe(
		ctx context.Context,
		data *recipemodel.RecipeCreate,
	) error
	DeleteRecipe(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type UpdateRecipeDetailStore interface {
	UpdateRecipeDetail(
		ctx context.Context,
		idRecipe string,
		idIngredient string,
		data *recipedetailmodel.RecipeDetailUpdate,
	) error
	CreateListRecipeDetail(
		ctx context.Context,
		data []recipedetailmodel.RecipeDetailCreate,
	) error
	FindListRecipeDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]recipedetailmodel.RecipeDetail, error)
	DeleteRecipeDetail(
		ctx context.Context,
		conditions map[string]interface{},
	) error
}

type updateFoodRepo struct {
	foodStore         UpdateFoodStore
	categoryFoodStore CreateOrDeleteCategoryFoodStore
	categoryStore     UpdateCategoryStore
	sizeFoodStore     UpdateSizeFoodStore
	recipeStore       UpdateRecipeStore
	recipeDetailStore UpdateRecipeDetailStore
}

func NewUpdateFoodRepo(
	foodStore UpdateFoodStore,
	categoryFoodStore CreateOrDeleteCategoryFoodStore,
	categoryStore UpdateCategoryStore,
	sizeFoodStore UpdateSizeFoodStore,
	recipeStore UpdateRecipeStore,
	recipeDetailStore UpdateRecipeDetailStore) *updateFoodRepo {
	return &updateFoodRepo{
		foodStore:         foodStore,
		categoryFoodStore: categoryFoodStore,
		categoryStore:     categoryStore,
		sizeFoodStore:     sizeFoodStore,
		recipeStore:       recipeStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (repo *updateFoodRepo) FindFood(
	ctx context.Context,
	id string) (*productmodel.Food, error) {
	currentFood, err := repo.foodStore.FindFood(
		ctx,
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return nil, err
	}
	return currentFood, nil
}

func (repo *updateFoodRepo) FindCategories(
	ctx context.Context,
	foodId string) ([]categorymodel.SimpleCategoryWithId, error) {
	simpleCategories, err := repo.categoryFoodStore.FindListCategories(ctx, foodId)
	if err != nil {
		return nil, err
	}
	return simpleCategories, nil
}

func (repo *updateFoodRepo) HandleCategory(
	ctx context.Context,
	foodId string,
	deletedCategoryFood []categorymodel.SimpleCategoryWithId,
	createdCategoryFood []categorymodel.SimpleCategoryWithId) error {
	for _, v := range deletedCategoryFood {
		if err := repo.updateAmountProduct(ctx, v.CategoryId, -1); err != nil {
			return err
		}
	}
	for _, v := range createdCategoryFood {
		if err := repo.updateAmountProduct(ctx, v.CategoryId, 1); err != nil {
			return err
		}
	}

	if err := repo.deleteCategoryFood(ctx, foodId, deletedCategoryFood); err != nil {
		return err
	}
	if err := repo.creatCategoryFood(ctx, foodId, createdCategoryFood); err != nil {
		return err
	}
	return nil
}

func (repo *updateFoodRepo) deleteCategoryFood(
	ctx context.Context,
	foodId string,
	deletedCategoryFood []categorymodel.SimpleCategoryWithId) error {
	for _, v := range deletedCategoryFood {
		if err := repo.categoryFoodStore.DeleteCategoryFood(ctx, map[string]interface{}{
			"foodId":     foodId,
			"categoryId": v.CategoryId,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) creatCategoryFood(
	ctx context.Context,
	foodId string,
	createdRecipeDetails []categorymodel.SimpleCategoryWithId) error {
	for _, v := range createdRecipeDetails {
		categoryFoodCreate := categoryfoodmodel.CategoryFoodCreate{
			FoodId:     foodId,
			CategoryId: v.CategoryId,
		}
		if err := repo.categoryFoodStore.CreateCategoryFood(
			ctx,
			&categoryFoodCreate); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) updateAmountProduct(
	ctx context.Context,
	categoryId string,
	amount int) error {
	updateModel := categorymodel.CategoryUpdateAmountProduct{AmountProduct: &amount}
	if err := repo.categoryStore.UpdateAmountProductCategory(
		ctx,
		categoryId,
		&updateModel); err != nil {
		return err
	}
	return nil
}

func (repo *updateFoodRepo) FindSizeFoods(
	ctx context.Context,
	foodId string) ([]sizefoodmodel.SizeFood, error) {
	sizeFoods, err := repo.sizeFoodStore.FindListSizeFood(ctx, foodId)
	if err != nil {
		return nil, err
	}
	return sizeFoods, nil
}

func (repo *updateFoodRepo) FindRecipeDetails(
	ctx context.Context,
	recipeId string) ([]recipedetailmodel.RecipeDetail, error) {
	recipeDetails, err := repo.recipeDetailStore.FindListRecipeDetail(
		ctx,
		map[string]interface{}{"recipeId": recipeId},
	)
	if err != nil {
		return nil, err
	}
	return recipeDetails, nil
}

func (repo *updateFoodRepo) HandleSizeFoods(
	ctx context.Context,
	foodId string,
	deletedSizeFood []sizefoodmodel.SizeFood,
	updatedSizeFood []sizefoodmodel.SizeFoodUpdate,
	mapDeletedRecipeDetails map[string][]recipedetailmodel.RecipeDetail,
	mapUpdatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailUpdate,
	mapCreatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailCreate,
	createdSizeFood []sizefoodmodel.SizeFoodCreate) error {
	if err := repo.handleDeleteSizeFoods(ctx, foodId, deletedSizeFood); err != nil {
		return err
	}
	if err := repo.handleUpdateSizeFoods(
		ctx,
		foodId,
		updatedSizeFood,
		mapDeletedRecipeDetails,
		mapUpdatedRecipeDetails,
		mapCreatedRecipeDetails); err != nil {
		return err
	}
	if err := repo.handleCreateSizeFoods(ctx, createdSizeFood); err != nil {
		return err
	}
	return nil
}

func (repo *updateFoodRepo) handleDeleteSizeFoods(
	ctx context.Context,
	foodId string,
	deletedSizeFood []sizefoodmodel.SizeFood) error {
	for _, v := range deletedSizeFood {
		if err := repo.recipeDetailStore.DeleteRecipeDetail(ctx, map[string]interface{}{
			"recipeId": v.RecipeId,
		}); err != nil {
			return err
		}

		if err := repo.sizeFoodStore.DeleteSizeFood(ctx, map[string]interface{}{
			"foodId": foodId,
			"sizeId": v.SizeId,
		}); err != nil {
			return err
		}

		if err := repo.recipeStore.DeleteRecipe(ctx, map[string]interface{}{
			"id": v.RecipeId,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) handleUpdateSizeFoods(
	ctx context.Context,
	foodId string,
	updatedSizeFood []sizefoodmodel.SizeFoodUpdate,
	mapDeletedRecipeDetails map[string][]recipedetailmodel.RecipeDetail,
	mapUpdatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailUpdate,
	mapCreatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailCreate) error {
	for _, v := range updatedSizeFood {
		if err := repo.sizeFoodStore.UpdateSizeFood(
			ctx,
			foodId,
			*v.SizeId,
			&v); err != nil {
			return err
		}

		if err := repo.handleUpdateRecipeDetail(
			ctx,
			*v.RecipeId,
			mapDeletedRecipeDetails[*v.RecipeId],
			mapUpdatedRecipeDetails[*v.RecipeId],
			mapCreatedRecipeDetails[*v.RecipeId],
		); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) handleUpdateRecipeDetail(
	ctx context.Context,
	recipeId string,
	deletedRecipeDetails []recipedetailmodel.RecipeDetail,
	updatedRecipeDetails []recipedetailmodel.RecipeDetailUpdate,
	createdRecipeDetails []recipedetailmodel.RecipeDetailCreate) error {
	if err := repo.deleteRecipeDetails(ctx, recipeId, deletedRecipeDetails); err != nil {
		return err
	}
	if err := repo.updateRecipeDetails(ctx, recipeId, updatedRecipeDetails); err != nil {
		return err
	}
	if err := repo.creatRecipeDetails(ctx, createdRecipeDetails); err != nil {
		return err
	}
	return nil
}

func (repo *updateFoodRepo) deleteRecipeDetails(
	ctx context.Context,
	recipeId string,
	deletedRecipeDetails []recipedetailmodel.RecipeDetail) error {
	for _, v := range deletedRecipeDetails {
		if err := repo.recipeDetailStore.DeleteRecipeDetail(ctx, map[string]interface{}{
			"recipeId":     recipeId,
			"ingredientId": v.IngredientId,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) updateRecipeDetails(
	ctx context.Context,
	recipeId string,
	updatedRecipeDetails []recipedetailmodel.RecipeDetailUpdate) error {
	for _, v := range updatedRecipeDetails {
		if err := repo.recipeDetailStore.UpdateRecipeDetail(
			ctx,
			recipeId,
			v.IngredientId,
			&v,
		); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) creatRecipeDetails(
	ctx context.Context,
	createdRecipeDetails []recipedetailmodel.RecipeDetailCreate) error {
	if err := repo.recipeDetailStore.CreateListRecipeDetail(
		ctx,
		createdRecipeDetails); err != nil {
		return err
	}
	return nil
}

func (repo *updateFoodRepo) handleCreateSizeFoods(
	ctx context.Context,
	createdSizeFood []sizefoodmodel.SizeFoodCreate) error {
	for _, v := range createdSizeFood {
		if err := repo.recipeStore.CreateRecipe(ctx, v.Recipe); err != nil {
			return err
		}

		if err := repo.recipeDetailStore.CreateListRecipeDetail(
			ctx,
			v.Recipe.Details); err != nil {
			return err
		}

		if err := repo.sizeFoodStore.CreateSizeFood(ctx, &v); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateFoodRepo) UpdateFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdateInfo) error {
	if err := repo.foodStore.UpdateFood(ctx, id, data); err != nil {
		return err
	}
	return nil
}
