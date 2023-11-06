package customermodel

import "coffee_shop_management_backend/common"

type CustomerCreate struct {
	Id    *string `json:"id" gorm:"column:id;"`
	Name  string  `json:"name" gorm:"column:name;"`
	Email string  `json:"email" gorm:"column:email;"`
	Phone string  `json:"phone" gorm:"column:phone;"`
}

func (*CustomerCreate) TableName() string {
	return common.TableCustomer
}

func (data *CustomerCreate) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrCustomerIdInvalid
	}
	if common.ValidateEmptyString(data.Name) {
		return ErrCustomerNameEmpty
	}
	if data.Email != "" && !common.ValidateEmail(data.Email) {
		return ErrCustomerEmailInvalid
	}
	if !common.ValidatePhone(data.Phone) {
		return ErrCustomerPhoneInvalid
	}
	return nil
}
