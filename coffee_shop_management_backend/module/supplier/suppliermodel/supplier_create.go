package suppliermodel

import "coffee_shop_management_backend/common"

type SupplierCreate struct {
	Id    *string `json:"id" gorm:"column:id;"`
	Name  string  `json:"name" gorm:"column:name;"`
	Email string  `json:"email" gorm:"column:email;"`
	Phone string  `json:"phone" gorm:"column:phone;"`
}

func (*SupplierCreate) TableName() string {
	return common.TableSupplier
}

func (data *SupplierCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrNameEmpty
	}
	if !common.ValidateEmail(data.Email) {
		return ErrEmailInvalid
	}
	if !common.ValidatePhone(data.Phone) {
		return ErrPhoneInvalid
	}
	return nil
}
