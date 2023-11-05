package categoryfoodmodel

import (
	"coffee_shop_management_backend/common"
)

type CategoryFoodCreate struct {
	FoodId     string `json:"foodId" gorm:"column:foodId;"`
	CategoryId string `json:"categoryId" gorm:"column:categoryId;"`
}

func (*CategoryFoodCreate) TableName() string {
	return common.TableCategoryFood
}

func (data *CategoryFoodCreate) Validate() *common.AppError {
	if !common.ValidateNotNilId(&data.FoodId) {
		return ErrIdFoodInvalid
	}
	if !common.ValidateNotNilId(&data.CategoryId) {
		return ErrIdCategoryInvalid
	}
	return nil
}
