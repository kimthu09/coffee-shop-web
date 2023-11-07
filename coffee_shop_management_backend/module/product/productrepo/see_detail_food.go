package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

type SeeDetailFoodStore interface {
	FindFood(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Food, error)
}

type seeDetailFoodRepo struct {
	store SeeDetailFoodStore
}

func NewSeeDetailFoodRepo(store SeeDetailFoodStore) *seeDetailFoodRepo {
	return &seeDetailFoodRepo{store: store}
}

func (repo *seeDetailFoodRepo) SeeDetailFood(
	ctx context.Context,
	foodId string) (*productmodel.Food, error) {
	result, err := repo.store.FindFood(
		ctx,
		map[string]interface{}{"id": foodId},
		"FoodSizes.Recipe.Details.Ingredient", "FoodCategories.Category")

	if err != nil {
		return nil, err
	}

	return result, nil
}
