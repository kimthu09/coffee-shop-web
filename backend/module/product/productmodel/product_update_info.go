package productmodel

import "coffee_shop_management_backend/common"

type ProductUpdateInfo struct {
	Name         *string `json:"name" gorm:"column:name;"`
	Description  *string `json:"description" gorm:"column:description;"`
	CookingGuide *string `json:"cookingGuide" gorm:"column:cookingGuide"`
}

func (data *ProductUpdateInfo) Validate() error {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrProductNameEmpty
	}
	return nil
}
