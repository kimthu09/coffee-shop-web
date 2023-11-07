package recipemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
)

type RecipeCreate struct {
	Id      string                                 `json:"-" gorm:"column:id;"`
	Details []recipedetailmodel.RecipeDetailCreate `json:"details" gorm:"-"`
}

func (*RecipeCreate) TableName() string {
	return common.TableRecipe
}

func (data *RecipeCreate) Validate() *common.AppError {
	if data.Details == nil || len(data.Details) == 0 {
		return ErrRecipeDetailsEmpty
	}
	mapAmountExist := make(map[string]int)
	for _, v := range data.Details {
		if err := v.Validate(); err != nil {
			return err
		}
		mapAmountExist[v.IngredientId]++
		if mapAmountExist[v.IngredientId] >= 2 {
			return ErrRecipeIngredientDuplicate
		}
	}
	return nil
}
