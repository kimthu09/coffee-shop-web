package recipemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
)

type RecipeUpdate struct {
	Details *[]recipedetailmodel.RecipeDetailUpdate `json:"details" gorm:"-"`
}

func (*RecipeUpdate) TableName() string {
	return common.TableRecipe
}

func (data *RecipeUpdate) Validate() *common.AppError {
	if data.Details != nil && len(*data.Details) == 0 {
		return ErrDetailsEmpty
	}
	mapAmountExist := make(map[string]int)
	for _, v := range *data.Details {
		if err := v.Validate(); err != nil {
			return err
		}
		if v.IngredientId != nil {
			mapAmountExist[*v.IngredientId]++
			if mapAmountExist[*v.IngredientId] >= 2 {
				return ErrDuplicateIngredient
			}
		}

	}
	return nil
}
