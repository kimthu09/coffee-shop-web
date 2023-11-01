package recipedetailmodel

import "coffee_shop_management_backend/common"

type RecipeDetailUpdate struct {
	IngredientId *string  `json:"ingredientId" gorm:"-"`
	AmountNeed   *float32 `json:"amountNeed" gorm:"column:amountNeed;"`
}

func (*RecipeDetailUpdate) TableName() string {
	return common.TableRecipeDetail
}

func (data *RecipeDetailUpdate) Validate() *common.AppError {
	if !common.ValidateNotNilId(data.IngredientId) {
		return ErrIngredientIdInvalid
	}
	if data.AmountNeed == nil {
		return ErrAmountNeedInvalid
	}
	if common.ValidateNotPositiveNumber(data.AmountNeed) {
		return ErrAmountNeedIsNotPositiveNumber
	}
	return nil
}
