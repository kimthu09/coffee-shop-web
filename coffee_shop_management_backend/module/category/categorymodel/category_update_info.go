package categorymodel

import (
	"coffee_shop_management_backend/common"
)

type CategoryUpdateInfo struct {
	Name        string  `json:"name" gorm:"column:name;"`
	Description *string `json:"description" gorm:"column:description;"`
}

func (*CategoryUpdateInfo) TableName() string {
	return common.TableCategory
}

func (c *CategoryUpdateInfo) Validate() error {
	if common.ValidateEmptyString(c.Name) {
		return ErrNameEmpty
	}

	return nil
}
