package ingredientdetailmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type IngredientDetail struct {
	IngredientId string  `json:"ingredientId" gorm:"column:ingredientId;"`
	ExpiryDate   string  `json:"expiryDate" gorm:"column:expiryDate;"`
	Amount       float32 `json:"amount" gorm:"column:amount;"`
}

func (*IngredientDetail) TableName() string {
	return common.TableIngredientDetail
}

var (
	ErrIngredientDetailViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view ingredient"),
	)
)
