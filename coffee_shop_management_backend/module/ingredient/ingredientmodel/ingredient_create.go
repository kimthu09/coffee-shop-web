package ingredientmodel

import (
	"coffee_shop_management_backend/common"
)

type IngredientCreate struct {
	Id          *string             `json:"id" gorm:"column:id;"`
	Name        string              `json:"name" gorm:"column:name;"`
	MeasureType *common.MeasureType `json:"measureType" gorm:"column:measureType;"`
	Price       float32             `json:"price" gorm:"column:price;"`
}

func (*IngredientCreate) TableName() string {
	return common.TableIngredient
}

func (data *IngredientCreate) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrIngredientNameEmpty
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrPriceIsNegativeNumber
	}
	if data.MeasureType == nil {
		return ErrMeasureTypeEmpty
	}
	return nil
}
