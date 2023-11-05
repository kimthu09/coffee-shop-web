package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"errors"
)

type Topping struct {
	*Product `json:",inline"`
	RecipeId string             `json:"-" gorm:"column:recipeId;"`
	Recipe   recipemodel.Recipe `json:"recipe" gorm:"-"`
}

func (*Topping) TableName() string {
	return common.TableTopping
}

var (
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
