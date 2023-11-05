package productmodel

import "coffee_shop_management_backend/common"

type ProductUpdate struct {
	Name         *string `json:"name" gorm:"column:name;"`
	Description  *string `json:"description" gorm:"column:description;"`
	CookingGuide *string `json:"cookingGuide" gorm:"column:cookingGuide"`
	IsActive     *bool   `json:"isActive" gorm:"column:isActive;"`
}

func (data *ProductUpdate) Validate() error {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrProductNameEmpty
	}
	return nil
}
