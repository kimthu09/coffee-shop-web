package productmodel

import (
	"coffee_shop_management_backend/common"
)

type ToppingUpdate struct {
	*ProductUpdate `json:",inline"`
}

func (*ToppingUpdate) TableName() string {
	return common.TableTopping
}

func (data *ToppingUpdate) Validate() error {
	if err := (*data.ProductUpdate).Validate(); err != nil {
		return err
	}
	return nil
}
