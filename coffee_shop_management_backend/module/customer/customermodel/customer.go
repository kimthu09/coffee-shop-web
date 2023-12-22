package customermodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"errors"
)

type Customer struct {
	Id       string                 `json:"id" gorm:"column:id;"`
	Name     string                 `json:"name" gorm:"column:name;"`
	Email    string                 `json:"email" gorm:"column:email;"`
	Phone    string                 `json:"phone" gorm:"column:phone;"`
	Point    float32                `json:"point" gorm:"column:point;"`
	Invoices []invoicemodel.Invoice `json:"invoices"`
}

func (*Customer) TableName() string {
	return common.TableCustomer
}

var (
	ErrCustomerIdInvalid = common.NewCustomError(
		errors.New("id of customer is invalid"),
		"id of customer is invalid",
		"ErrCustomerIdInvalid",
	)
	ErrCustomerNameEmpty = common.NewCustomError(
		errors.New("name of customer is empty"),
		"name of customer is empty",
		"ErrCustomerNameEmpty",
	)
	ErrCustomerPhoneInvalid = common.NewCustomError(
		errors.New("phone of customer is invalid"),
		"phone of customer is invalid",
		"ErrCustomerPhoneInvalid",
	)
	ErrCustomerEmailInvalid = common.NewCustomError(
		errors.New("email of customer is invalid"),
		"email of customer is invalid",
		"ErrCustomerEmailInvalid",
	)
	ErrCustomerDebtPayNotExist = common.NewCustomError(
		errors.New("debt pay is not exist"),
		"debt pay is not exist",
		"ErrCustomerDebtPayNotExist",
	)
	ErrCustomerDebtPayIsInvalid = common.NewCustomError(
		errors.New("debt pay is invalid"),
		"debt pay is invalid",
		"ErrCustomerDebtPayIsInvalid",
	)
	ErrCustomerIdDuplicate = common.ErrDuplicateKey(
		errors.New("id of customer is duplicate"),
	)
	ErrCustomerPhoneDuplicate = common.ErrDuplicateKey(
		errors.New("phone of customer is duplicate"),
	)
	ErrCustomerCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create customer"),
	)
	ErrCustomerUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update info customer"),
	)
	ErrCustomerViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view customer"),
	)
)
