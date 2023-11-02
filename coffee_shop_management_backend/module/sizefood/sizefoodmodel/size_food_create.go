package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type SizeFoodCreate struct {
	FoodId   string                    `json:"-" gorm:"column:foodId;"`
	SizeId   string                    `json:"-" gorm:"column:sizeId;"`
	Name     string                    `json:"name" gorm:"column:name"`
	Cost     float64                   `json:"cost" gorm:"column:cost"`
	Price    float64                   `json:"price" gorm:"column:price"`
	RecipeId string                    `json:"-" gorm:"column:recipeId;"`
	Recipe   *recipemodel.RecipeCreate `json:"recipe" gorm:"-"`
}

func (*SizeFoodCreate) TableName() string {
	return common.TableSizeFood
}

func (data *SizeFoodCreate) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrNameEmpty
	}
	if common.ValidateNegativeNumber(data.Cost) {
		return ErrCostIsNegativeNumber
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrPriceIsNegativeNumber
	}
	if data.Recipe == nil {
		return ErrRecipeEmpty
	}
	if err := data.Recipe.Validate(); err != nil {
		return err
	}
	return nil
}
