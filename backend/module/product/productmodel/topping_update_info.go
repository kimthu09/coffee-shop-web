package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type ToppingUpdateInfo struct {
	*ProductUpdateInfo `json:",inline"`
	Cost               *int                      `json:"cost" gorm:"column:cost;"`
	Price              *int                      `json:"price" gorm:"column:price;"`
	Recipe             *recipemodel.RecipeUpdate `json:"recipe" gorm:"-"`
}

func (*ToppingUpdateInfo) TableName() string {
	return common.TableTopping
}

func (data *ToppingUpdateInfo) Validate() error {
	if data.ProductUpdateInfo != nil {
		if err := (*data.ProductUpdateInfo).Validate(); err != nil {
			return err
		}
	}
	if data.Cost != nil && common.ValidateNegativeNumber(*data.Cost) {
		return ErrToppingCostIsNegativeNumber
	}
	if data.Price != nil && common.ValidateNegativeNumber(*data.Price) {
		return ErrToppingPriceIsNegativeNumber
	}
	if data.Recipe != nil {
		if err := data.Recipe.Validate(); err != nil {
			return err
		}
	}
	return nil
}
