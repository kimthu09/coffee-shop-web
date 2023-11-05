package ingredientdetailmodel

import "coffee_shop_management_backend/common"

type IngredientDetailUpdate struct {
	IngredientId string  `json:"-" gorm:"-"`
	ExpiryDate   string  `json:"-" gorm:"-"`
	Amount       float32 `json:"amount" gorm:"-"`
}

func (*IngredientDetailUpdate) TableName() string {
	return common.TableIngredientDetail
}
