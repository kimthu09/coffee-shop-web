package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"errors"
)

type Ingredient struct {
	Id          string            `json:"id" gorm:"column:id;"`
	Name        string            `json:"name" gorm:"column:name;"`
	TotalAmount float32           `json:"totalAmount" gorm:"column:totalAmount;"`
	MeasureType *enum.MeasureType `json:"measureType" gorm:"column:measureType;"`
	Price       float32           `json:"price" gorm:"column:price;"`
}

func (*Ingredient) TableName() string {
	return common.TableIngredient
}

var (
	ErrIngredientIdInvalid = common.NewCustomError(
		errors.New("id of ingredient is invalid"),
		"id of ingredient is invalid",
		"ErrIngredientIdInvalid",
	)
	ErrIngredientNameEmpty = common.NewCustomError(
		errors.New("name of ingredient is empty"),
		"name of ingredient is empty",
		"ErrIngredientNameEmpty",
	)
	ErrIngredientPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price of ingredient is negative number"),
		"price of ingredient is negative number",
		"ErrIngredientPriceIsNegativeNumber",
	)
	ErrIngredientMeasureTypeEmpty = common.NewCustomError(
		errors.New("measure type of ingredient is empty"),
		"measure type of ingredient is empty",
		"ErrIngredientMeasureTypeEmpty",
	)
	ErrIngredientAmountUpdateInvalid = common.NewCustomError(
		errors.New("amount need to update for the ingredient is invalid"),
		"amount need to update for the ingredient is invalid",
		"ErrIngredientAmountUpdateInvalid",
	)
	ErrIngredientIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of ingredient is duplicate"),
	)
	ErrIngredientNameDuplicate = common.ErrDuplicateKey(
		errors.New("name of ingredient is duplicate"),
	)
	ErrIngredientCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create import note"),
	)
)
