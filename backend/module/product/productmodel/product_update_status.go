package productmodel

import "coffee_shop_management_backend/common"

type ProductUpdateStatus struct {
	ProductId string `json:"id" gorm:"-"`
	IsActive  *bool  `json:"isActive" gorm:"column:isActive;"`
}

func (data *ProductUpdateStatus) Validate() error {
	if !common.ValidateNotNilId(&data.ProductId) {
		return ErrProductIdInvalid
	}
	if data.IsActive == nil {
		return ErrProductIsActiveEmpty
	}
	return nil
}
