package ingredientmodel

import "coffee_shop_management_backend/common"

type IngredientUpdate struct {
	Amount float32 `json:"amount" gorm:"-"`
}

func (*IngredientUpdate) TableName() string {
	return common.TableIngredient
}

func (data *IngredientUpdate) Validate() *common.AppError {
	return nil
}
