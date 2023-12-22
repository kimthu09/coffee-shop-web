package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"errors"
)

type Topping struct {
	*Product `json:",inline"`
	Cost     int                 `json:"cost" gorm:"column:cost;"`
	Price    int                 `json:"price" gorm:"column:price;"`
	RecipeId string              `json:"-" gorm:"column:recipeId;"`
	Recipe   *recipemodel.Recipe `json:"recipe" gorm:"foreignkey:recipeId"`
}

func (*Topping) TableName() string {
	return common.TableTopping
}

var (
	ErrToppingProductInfoEmpty = common.NewCustomError(
		errors.New("product info of topping is empty"),
		"product info of topping is empty",
		"ErrToppingProductInfoEmpty",
	)
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
	ErrToppingRecipeEmpty = common.NewCustomError(
		errors.New("recipe is empty"),
		"recipe is empty",
		"ErrToppingRecipeEmpty",
	)
	ErrToppingCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create topping"),
	)
	ErrToppingUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info topping"),
	)
	ErrToppingChangeStatusNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to change status topping"),
	)
	ErrToppingViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view topping"),
	)
)
