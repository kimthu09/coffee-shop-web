package recipedetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type RecipeDetail struct {
	RecipeId     string  `json:"recipeId" gorm:"column:recipeId;"`
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	AmountNeed   float32 `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetail) TableName() string {
	return common.TableRecipeDetail
}

var (
	ErrRecipeDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"id of ingredient is invalid",
		"ErrRecipeDetailIngredientIdInvalid",
	)
	ErrRecipeDetailAmountNeedIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount need is not positive number"),
		"amount need is not positive number",
		"ErrRecipeDetailAmountNeedIsNotPositiveNumber",
	)
)
