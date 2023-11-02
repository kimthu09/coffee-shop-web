package productmodel

import "coffee_shop_management_backend/common"

type ProductCreate struct {
	Id          *string `json:"id" gorm:"column:id;"`
	Name        string  `json:"name" gorm:"column:name;"`
	Description string  `json:"description" gorm:"column:description;"`
}

func (data *ProductCreate) Validate() error {
	if !common.ValidateId(data.Id) {
		return ErrIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrNameEmpty
	}
	return nil
}
