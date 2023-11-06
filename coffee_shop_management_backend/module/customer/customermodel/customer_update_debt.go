package customermodel

import "coffee_shop_management_backend/common"

type CustomerUpdateDebt struct {
	Amount *float32 `json:"amount" gorm:"-"`
}

func (*CustomerUpdateDebt) TableName() string {
	return common.TableCustomer
}

func (data *CustomerUpdateDebt) Validate() *common.AppError {
	if data.Amount == nil {
		return ErrCustomerDebtPayNotExist
	}
	if *data.Amount == 0 {
		return ErrCustomerDebtPayIsInvalid
	}
	return nil
}
