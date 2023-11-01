package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type SizeFood struct {
	FoodId   string  `json:"foodId" gorm:"column:foodId;"`
	SizeId   string  `json:"sizeId" gorm:"column:sizeId;"`
	Name     string  `json:"name" gorm:"column:name;"`
	Cost     float64 `json:"cost" gorm:"column:cost;"`
	Price    float64 `json:"price" gorm:"column:price;"`
	RecipeId string  `json:"recipeId" gorm:"column:recipeId;"`
}

func (*SizeFood) TableName() string {
	return common.TableSizeFood
}

var (
	ErrNameEmpty = common.NewCustomError(
		errors.New("name of size is empty"),
		"name of size is empty",
		"ErrNameEmpty",
	)
	ErrCostIsNegativeNumber = common.NewCustomError(
		errors.New("cost is negative number"),
		"cost is negative number",
		"ErrCostIsNegativeNumber",
	)
	ErrPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price is negative number"),
		"price is negative number",
		"ErrPriceIsNegativeNumber",
	)
	ErrRecipeEmpty = common.NewCustomError(
		errors.New("recipe is empty"),
		"recipe is empty",
		"ErrRecipeEmpty",
	)
)
