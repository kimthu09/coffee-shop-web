package suppliermodel

import "coffee_shop_management_backend/common"

type SupplierUpdateDebt struct {
	Amount *float32 `json:"amount" gorm:"-"`
}

func (*SupplierUpdateDebt) TableName() string {
	return common.TableSupplier
}

func (data *SupplierUpdateDebt) Validate() *common.AppError {
	if data.Amount == nil {
		return ErrDebtPayNotExist
	}
	return nil
}
