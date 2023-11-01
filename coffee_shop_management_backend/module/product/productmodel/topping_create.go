package productmodel

import (
	"coffee_shop_management_backend/common"
)

type ToppingCreate struct {
	*ProductCreate `json:",inline"`
}

func (*ToppingCreate) TableName() string {
	return common.TableTopping
}

func (data *ToppingCreate) Validate() error {
	if err := (*data.ProductCreate).Validate(); err != nil {
		return err
	}
	return nil
}
