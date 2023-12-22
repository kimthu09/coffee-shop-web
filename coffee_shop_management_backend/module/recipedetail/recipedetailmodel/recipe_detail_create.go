package recipedetailmodel

import "coffee_shop_management_backend/common"

type RecipeDetailCreate struct {
	RecipeId     string `json:"-" gorm:"column:recipeId;"`
	IngredientId string `json:"ingredientId" gorm:"column:ingredientId;"`
	AmountNeed   int    `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetailCreate) TableName() string {
	return common.TableRecipeDetail
}

func (data *RecipeDetailCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrRecipeDetailIngredientIdInvalid
	}
	if common.ValidateNotPositiveNumber(data.AmountNeed) {
		return ErrRecipeDetailAmountNeedIsNotPositiveNumber
	}
	return nil
}
