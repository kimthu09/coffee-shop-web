package recipemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"errors"
)

type Recipe struct {
	Id      string                            `json:"id" gorm:"column:id;"`
	Details *[]recipedetailmodel.RecipeDetail `json:"details" gorm:"-"`
}

func (*Recipe) TableName() string {
	return common.TableRecipe
}

var (
	ErrRecipeIngredientDuplicate = common.NewCustomError(
		errors.New("ingredient for recipe is duplicate"),
		"ingredient for recipe is duplicate",
		"ErrRecipeIngredientDuplicate",
	)
	ErrRecipeDetailsEmpty = common.NewCustomError(
		errors.New("ingredient for recipe is empty"),
		"ingredient for recipe is empty",
		"ErrRecipeDetailsEmpty",
	)
)
