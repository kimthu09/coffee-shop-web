package generator

import "coffee_shop_management_backend/common"

type IdGenerator interface {
	GenerateId() (string, error)
	IdProcess(id *string) (*string, *common.AppError)
}
