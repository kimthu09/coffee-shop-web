package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type ToppingUpdate struct {
	*ProductUpdate `json:",inline"`
	Recipe         *recipemodel.RecipeUpdate `json:"recipe" gorm:"-"`
}

func (*ToppingUpdate) TableName() string {
	return common.TableTopping
}

func (data *ToppingUpdate) Validate() error {
	if err := (*data.ProductUpdate).Validate(); err != nil {
		return err
	}
	if data.Recipe != nil {
		if err := data.Recipe.Validate(); err != nil {
			return err
		}
	}
	return nil
}
