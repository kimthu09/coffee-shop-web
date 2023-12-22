package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
)

type SimpleIngredient struct {
	Id          string            `json:"id" gorm:"column:id;"`
	Name        string            `json:"name" gorm:"column:name;"`
	MeasureType *enum.MeasureType `json:"measureType" gorm:"column:measureType;"`
}

func (*SimpleIngredient) TableName() string {
	return common.TableIngredient
}
