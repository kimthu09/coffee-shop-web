package productmodel

import (
	"coffee_shop_management_backend/common"
)

type ToppingUpdateStatus struct {
	*ProductUpdateStatus `json:",inline"`
}

func (*ToppingUpdateStatus) TableName() string {
	return common.TableTopping
}

func (data *ToppingUpdateStatus) Validate() error {
	if err := (*data.ProductUpdateStatus).Validate(); err != nil {
		return err
	}
	return nil
}
