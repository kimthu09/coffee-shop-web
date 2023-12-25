package recipedetailmodel

import "coffee_shop_management_backend/common"

type RecipeDetailUpdate struct {
	IngredientId string `json:"ingredientId" gorm:"-"`
	AmountNeed   int    `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetailUpdate) TableName() string {
	return common.TableRecipeDetail
}

func (data *RecipeDetailUpdate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.IngredientId) {
		return ErrRecipeDetailIngredientIdInvalid
	}
	if common.ValidateNotPositiveNumber(data.AmountNeed) {
		return ErrRecipeDetailAmountNeedIsNotPositiveNumber
	}
	return nil
}
