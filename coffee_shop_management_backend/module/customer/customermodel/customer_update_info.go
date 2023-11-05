package customermodel

import "coffee_shop_management_backend/common"

type CustomerUpdateInfo struct {
	Name  *string `json:"name" gorm:"column:name;"`
	Email *string `json:"email" gorm:"column:email;"`
	Phone *string `json:"phone" gorm:"column:phone;"`
}

func (*CustomerUpdateInfo) TableName() string {
	return common.TableCustomer
}

func (data *CustomerUpdateInfo) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrCustomerNameEmpty
	}
	if data.Email != nil && *data.Email != "" && !common.ValidateEmail(*data.Email) {
		return ErrCustomerEmailInvalid
	}
	if data.Phone != nil && !common.ValidatePhone(*data.Phone) {
		return ErrCustomerPhoneInvalid
	}
	return nil
}
