package suppliermodel

import "coffee_shop_management_backend/common"

type SupplierUpdateDebt struct {
	Id        *string `json:"id" gorm:"-"`
	Amount    *int    `json:"amount" gorm:"-"`
	CreatedBy string  `json:"-" gorm:"-"`
}

func (*SupplierUpdateDebt) TableName() string {
	return common.TableSupplier
}

func (data *SupplierUpdateDebt) Validate() *common.AppError {
	if !common.ValidateId(data.Id) {
		return ErrSupplierDebtIdInvalid
	}
	if data.Amount == nil {
		return ErrDebtPayNotExist
	}
	if *data.Amount == 0 {
		return ErrDebtPayIsInvalid
	}
	return nil
}
