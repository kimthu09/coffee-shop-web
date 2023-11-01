package productmodel

import "coffee_shop_management_backend/common"

type ProductUpdate struct {
	Name         *string `json:"name" gorm:"column:name;"`
	Description  *string `json:"description" gorm:"column:description;"`
	ProductSizes *Sizes  `json:"sizes" gorm:"column:sizes;"`
	IsActive     *bool   `json:"isActive" gorm:"column:isActive;"`
}

func (data *ProductUpdate) Validate() error {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrNameEmpty
	}
	if data.ProductSizes != nil && len(*data.ProductSizes) != 0 {
		return ErrSizeNotExist
	}
	return nil
}
