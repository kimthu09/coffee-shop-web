package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type SizeFoodUpdate struct {
	SizeId   *string                   `json:"sizeId" gorm:"-"`
	Name     *string                   `json:"name" gorm:"column:name"`
	Cost     *int                      `json:"cost" gorm:"column:cost"`
	Price    *int                      `json:"price" gorm:"column:price"`
	RecipeId *string                   `json:"-" gorm:"-"`
	Recipe   *recipemodel.RecipeUpdate `json:"recipe" gorm:"-"`
}

func (*SizeFoodUpdate) TableName() string {
	return common.TableSizeFood
}

func (data *SizeFoodUpdate) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrSizeFoodNameEmpty
	}
	if data.Cost != nil && common.ValidateNegativeNumberInt(*data.Cost) {
		return ErrSizeFoodCostIsNegativeNumber
	}
	if data.Price != nil && common.ValidateNegativeNumberInt(*data.Price) {
		return ErrSizeFoodPriceIsNegativeNumber
	}
	if data.Recipe != nil {
		if err := data.Recipe.Validate(); err != nil {
			return err
		}
	}
	return nil
}
