package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"errors"
)

type Topping struct {
	*Product `json:",inline"`
	Cost     float32            `json:"cost" gorm:"column:cost;"`
	Price    float32            `json:"price" gorm:"column:price;"`
	RecipeId string             `json:"-" gorm:"column:recipeId;"`
	Recipe   recipemodel.Recipe `json:"recipe" gorm:"-"`
}

func (*Topping) TableName() string {
	return common.TableTopping
}

var (
	ErrToppingCostIsNegativeNumber = common.NewCustomError(
		errors.New("cost is negative number"),
		"cost is negative number",
		"ErrSizeFoodCostIsNegativeNumber",
	)
	ErrToppingPriceIsNegativeNumber = common.NewCustomError(
		errors.New("price is negative number"),
		"price is negative number",
		"ErrSizeFoodPriceIsNegativeNumber",
	)
	ErrToppingIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of topping is duplicate"),
	)
	ErrToppingNameDuplicate = common.ErrDuplicateKey(
		errors.New("name of topping is duplicate"),
	)
	ErrRecipeEmpty = common.NewCustomError(
		errors.New("recipe is empty"),
		"recipe is empty",
		"ErrRecipeEmpty",
	)
)
