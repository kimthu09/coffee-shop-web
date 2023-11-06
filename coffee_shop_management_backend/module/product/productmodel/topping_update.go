package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type ToppingUpdate struct {
	*ProductUpdate `json:",inline"`
	Cost           *float32                  `json:"cost" gorm:"column:cost;"`
	Price          *float32                  `json:"price" gorm:"column:price;"`
	Recipe         *recipemodel.RecipeUpdate `json:"recipe" gorm:"-"`
}

func (*ToppingUpdate) TableName() string {
	return common.TableTopping
}

func (data *ToppingUpdate) Validate() error {
	if err := (*data.ProductUpdate).Validate(); err != nil {
		return err
	}
	if data.Cost != nil && common.ValidateNegativeNumber(data.Cost) {
		return ErrToppingCostIsNegativeNumber
	}
	if data.Price != nil && common.ValidateNegativeNumber(data.Price) {
		return ErrToppingPriceIsNegativeNumber
	}
	if data.Recipe != nil {
		if err := data.Recipe.Validate(); err != nil {
			return err
		}
	}
	return nil
}
