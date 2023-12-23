package recipedetailmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"errors"
)

type RecipeDetail struct {
	RecipeId     string                           `json:"recipeId" gorm:"column:recipeId;"`
	IngredientId string                           `json:"-" gorm:"column:ingredientId;"`
	Ingredient   ingredientmodel.SimpleIngredient `json:"ingredient" gorm:"foreignKey:IngredientId;references:Id"`
	AmountNeed   int                              `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetail) TableName() string {
	return common.TableRecipeDetail
}

var (
	ErrRecipeDetailIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"Mã của nguyên vật liệu không hợp lệ",
		"ErrRecipeDetailIngredientIdInvalid",
	)
	ErrRecipeDetailAmountNeedIsNotPositiveNumber = common.NewCustomError(
		errors.New("amount need is not positive number"),
		"Lượng nguyên vật liệu đang không phải là số dương",
		"ErrRecipeDetailAmountNeedIsNotPositiveNumber",
	)
)
