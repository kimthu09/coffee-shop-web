package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
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

type updateToppingRepo struct {
	toppingStore      UpdateToppingStore
	ingredientStore   CheckIngredientStore
	recipeDetailStore UpdateRecipeDetailStore
}

func NewUpdateToppingRepo(
	toppingStore UpdateToppingStore,
	ingredientStore CheckIngredientStore,
	recipeDetailStore UpdateRecipeDetailStore) *updateToppingRepo {
	return &updateToppingRepo{
		toppingStore:      toppingStore,
		ingredientStore:   ingredientStore,
		recipeDetailStore: recipeDetailStore,
	}
}

func (repo *updateToppingRepo) CheckIngredient(
	ctx context.Context,
	data []recipedetailmodel.RecipeDetailUpdate) error {
	for _, recipeDetail := range data {
		if _, err := repo.ingredientStore.FindIngredient(
			ctx,
			map[string]interface{}{"id": recipeDetail.IngredientId},
		); err != nil {
			return err
		}
	}
	return nil
}

func (repo *updateToppingRepo) FindTopping(
	ctx context.Context,
	id string) (*productmodel.Topping, error) {
	currentTopping, err := repo.toppingStore.FindTopping(
		ctx,
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return nil, err
	}
	return currentTopping, nil
}

func (repo *updateToppingRepo) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdate) error {
	if err := repo.toppingStore.UpdateTopping(ctx, id, data); err != nil {
		return err
	}
	return nil
}

func (repo *updateToppingRepo) FindRecipeDetails(
	ctx context.Context,
	recipeId string) ([]recipedetailmodel.RecipeDetail, error) {
	currentRecipeDetails, err := repo.recipeDetailStore.FindListRecipeDetail(
		ctx,
		map[string]interface{}{"recipeId": recipeId},
	)
	if err != nil {
		return nil, err
	}
	return currentRecipeDetails, nil
}

func (repo *updateToppingRepo) UpdateRecipeDetailsOfRecipe(
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

func (repo *updateToppingRepo) deleteRecipeDetails(
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

func (repo *updateToppingRepo) updateRecipeDetails(
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

func (repo *updateToppingRepo) creatRecipeDetails(
	ctx context.Context,
	createdRecipeDetails []recipedetailmodel.RecipeDetailCreate) error {
	if err := repo.recipeDetailStore.CreateListRecipeDetail(
		ctx,
		createdRecipeDetails); err != nil {
		return err
	}
	return nil
}
