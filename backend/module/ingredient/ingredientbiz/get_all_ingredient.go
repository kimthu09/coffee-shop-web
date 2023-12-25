package ingredientbiz

import (
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

type GetAllIngredientStore interface {
	GetAllIngredient(
		ctx context.Context) ([]ingredientmodel.Ingredient, error)
}

type getAllIngredientBiz struct {
	store GetAllIngredientStore
}

func NewGetAllIngredientBiz(
	store GetAllIngredientStore) *getAllIngredientBiz {
	return &getAllIngredientBiz{store: store}
}

func (biz *getAllIngredientBiz) GetAllIngredient(
	ctx context.Context) ([]ingredientmodel.Ingredient, error) {
	result, err := biz.store.GetAllIngredient(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
