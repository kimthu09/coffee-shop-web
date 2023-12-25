package suppliermodel

import "coffee_shop_management_backend/common"

type SupplierUpdateInfo struct {
	Name  *string `json:"name" gorm:"column:name;"`
	Email *string `json:"email" gorm:"column:email;"`
	Phone *string `json:"phone" gorm:"column:phone;"`
}

func (*SupplierUpdateInfo) TableName() string {
	return common.TableSupplier
}

func (data *SupplierUpdateInfo) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrSupplierNameEmpty
	}
	if data.Email != nil && *data.Email != "" && !common.ValidateEmail(*data.Email) {
		return ErrSupplierEmailInvalid
	}
	if data.Phone != nil && !common.ValidatePhone(*data.Phone) {
		return ErrSupplierPhoneInvalid
	}
	return nil
}
