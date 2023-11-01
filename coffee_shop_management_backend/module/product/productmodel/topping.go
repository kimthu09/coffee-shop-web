package productmodel

import (
	"coffee_shop_management_backend/common"
)

type Topping struct {
	*Product `json:",inline"`
}

func (*Topping) TableName() string {
	return common.TableTopping
}
