package recipemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"errors"
)

type Recipe struct {
	Id      string        `json:"id" gorm:"column:id;"`
	Details RecipeDetails `json:"details"`
}

func (*Recipe) TableName() string {
	return common.TableRecipe
}

type RecipeDetails []recipedetailmodel.RecipeDetail

func (*RecipeDetails) TableName() string {
	return common.TableRecipeDetail
}

var (
	ErrRecipeIngredientDuplicate = common.NewCustomError(
		errors.New("ingredient for recipe is duplicate"),
		"Tồn tại 2 nguyên vật liệu trùng nhau trong công thức",
		"ErrRecipeIngredientDuplicate",
	)
	ErrRecipeDetailsEmpty = common.NewCustomError(
		errors.New("ingredient for recipe is empty"),
		"Công thức đang không có nguyên vật liệu",
		"ErrRecipeDetailsEmpty",
	)
)
