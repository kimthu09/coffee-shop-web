package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type SizeFood struct {
	FoodId   string  `json:"foodId" gorm:"column:foodId;"`
	SizeId   string  `json:"sizeId" gorm:"column:sizeId;"`
	Name     string  `json:"name" gorm:"column:name;"`
	Cost     float32 `json:"cost" gorm:"column:cost;"`
	Price    float32 `json:"price" gorm:"column:price;"`
	RecipeId string  `json:"recipeId" gorm:"column:recipeId;"`
}

func (*SizeFood) TableName() string {
	return common.TableSizeFood
}

var (
	ErrSizeFoodNameEmpty = common.NewCustomError(
		errors.New("name of size is empty"),
		"name of size is empty",
		"ErrSizeFoodNameEmpty",
	)
	ErrSizeFoodCostIsNegativeNumber = common.NewCustomError(
		errors.New("cost is negative number"),
		"cost is negative number",
		"ErrSizeFoodCostIsNegativeNumber",
	)
	ErrSizeFoodPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price is negative number"),
		"price is negative number",
		"ErrSizeFoodPriceIsNegativeNumber",
	)
	ErrSizeFoodRecipeEmpty = common.NewCustomError(
		errors.New("recipe is empty"),
		"recipe is empty",
		"ErrSizeFoodRecipeEmpty",
	)
)
