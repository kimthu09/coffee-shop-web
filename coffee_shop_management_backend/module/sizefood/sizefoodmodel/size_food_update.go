package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
)

type SizeFoodUpdate struct {
	SizeId   *string                   `json:"sizeId" gorm:"-"`
	Name     *string                   `json:"name" gorm:"column:name"`
	Cost     *float32                  `json:"cost" gorm:"column:cost"`
	Price    *float32                  `json:"price" gorm:"column:price"`
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
	if data.Cost != nil && common.ValidateNegativeNumber(data.Cost) {
		return ErrSizeFoodCostIsNegativeNumber
	}
	if data.Price != nil && common.ValidateNegativeNumber(data.Price) {
		return ErrSizeFoodPriceIsNegativeNumber
	}
	if data.Recipe != nil {
		if err := data.Recipe.Validate(); err != nil {
			return err
		}
	}
	return nil
}
