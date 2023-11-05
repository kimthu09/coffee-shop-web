package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type ToppingCreate struct {
	*ProductCreate `json:",inline"`
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
	if data.Recipe == nil {
		return ErrRecipeEmpty
	}
	if err := data.Recipe.Validate(); err != nil {
		return err
	}
	return nil
}
