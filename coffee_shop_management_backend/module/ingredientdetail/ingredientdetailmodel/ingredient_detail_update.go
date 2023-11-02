package ingredientdetailmodel

import "coffee_shop_management_backend/common"

type IngredientDetailUpdate struct {
	Amount float32 `json:"amount" gorm:"-"`
}

func (*IngredientDetailUpdate) TableName() string {
	return common.TableIngredientDetail
}
