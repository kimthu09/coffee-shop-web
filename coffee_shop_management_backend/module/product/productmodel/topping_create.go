package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type ToppingCreate struct {
	*ProductCreate `json:",inline"`
	Cost           float32                   `json:"cost" gorm:"column:cost;"`
	Price          float32                   `json:"price" gorm:"column:price;"`
	RecipeId       string                    `json:"-" gorm:"column:recipeId;"`
	Recipe         *recipemodel.RecipeCreate `json:"recipe" gorm:"-"`
}

func (*ToppingCreate) TableName() string {
	return common.TableTopping
}

func (data *ToppingCreate) Validate() error {
	if err := (*data.ProductCreate).Validate(); err != nil {
		return err
	}
	if common.ValidateNegativeNumber(data.Cost) {
		return ErrToppingCostIsNegativeNumber
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrToppingPriceIsNegativeNumber
	}
	if data.Recipe == nil {
		return ErrRecipeEmpty
	}
	if err := data.Recipe.Validate(); err != nil {
		return err
	}
	return nil
}
