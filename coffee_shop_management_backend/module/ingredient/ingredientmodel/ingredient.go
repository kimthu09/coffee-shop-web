package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Ingredient struct {
	Id          string              `json:"id" gorm:"column:id;"`
	Name        string              `json:"name" gorm:"column:name;"`
	TotalAmount float32             `json:"totalAmount" gorm:"column:totalAmount;"`
	MeasureType *common.MeasureType `json:"measureType" gorm:"column:measureType;"`
	Price       float32             `json:"price" gorm:"column:price;"`
}

func (*Ingredient) TableName() string {
	return common.TableIngredient
}

var (
	ErrIngredientNameEmpty = common.NewCustomError(
		errors.New("name of ingredient is empty"),
		"name of ingredient is empty",
		"ErrIngredientNameEmpty",
	)
	ErrPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price is negative number"),
		"price is negative number",
		"ErrPriceIsNegativeNumber",
	)
	ErrMeasureTypeEmpty = common.NewCustomError(
		errors.New("measure type of ingredient is empty"),
		"measure type of ingredient is empty",
		"ErrMeasureTypeEmpty",
	)
)
