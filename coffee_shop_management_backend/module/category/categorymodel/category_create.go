package categorymodel

import (
	"coffee_shop_management_backend/common"
)

type CategoryCreate struct {
	Id          string `json:"-" gorm:"column:id;"`
	Name        string `json:"name" gorm:"column:name;"`
	Description string `json:"description" gorm:"column:description;"`
}

func (*CategoryCreate) TableName() string {
	return common.TableCategory
}

func (data *CategoryCreate) Validate() error {
	if common.ValidateEmptyString(data.Name) {
		return ErrCategoryNameEmpty
	}
	return nil
}
