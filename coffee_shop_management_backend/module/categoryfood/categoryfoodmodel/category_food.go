package categoryfoodmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type CategoryFood struct {
	FoodId     string `json:"foodId" gorm:"column:foodId;"`
	CategoryId string `json:"categoryId" gorm:"column:categoryId;"`
}

func (*CategoryFood) TableName() string {
	return common.TableCategoryFood
}

var (
	ErrIdFoodInvalid = common.NewCustomError(
		errors.New("id of food is invalid"),
		"id of food is invalid",
		"ErrIdFoodInvalid",
	)
	ErrIdCategoryInvalid = common.NewCustomError(
		errors.New("id of category is invalid"),
		"id of category is invalid",
		"ErrIdCategoryInvalid",
	)
)
