package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
)

type IngredientCreate struct {
	Id          *string           `json:"id" gorm:"column:id;"`
	Name        string            `json:"name" gorm:"column:name;"`
	MeasureType *enum.MeasureType `json:"measureType" gorm:"column:measureType;"`
	Price       float32           `json:"price" gorm:"column:price;"`
}

func (*IngredientCreate) TableName() string {
	return common.TableIngredient
}

func (data *IngredientCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrIngredientIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrIngredientNameEmpty
	}
	if data.MeasureType == nil {
		return ErrIngredientMeasureTypeEmpty
	}
	if common.ValidateNegativeNumber(data.Price) {
		return ErrIngredientPriceIsNegativeNumber
	}
	return nil
}

func (data *IngredientCreate) Round() {
	common.CustomRound(&data.Price)
}
