package ingredientmodel

import "coffee_shop_management_backend/common"

type IngredientUpdateAmount struct {
	Amount float32 `json:"amount" gorm:"-"`
}

func (*IngredientUpdateAmount) TableName() string {
	return common.TableIngredient
}

func (data *IngredientUpdateAmount) Validate() *common.AppError {
	if data.Amount == 0 {
		return ErrIngredientAmountUpdateInvalid
	}
	return nil
}
