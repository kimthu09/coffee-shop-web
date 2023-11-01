package ingredientdetailmodel

import "coffee_shop_management_backend/common"

type IngredientDetailCreate struct {
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string  `json:"expiryDate" gorm:"column:expiryDate;"`
	Amount       float32 `json:"-" gorm:"column:amount;"`
}

func (*IngredientDetailCreate) TableName() string {
	return common.TableIngredientDetail
}
