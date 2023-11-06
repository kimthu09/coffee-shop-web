package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) CreateTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY", "name"); key {
			case "PRIMARY":
				return productmodel.ErrToppingIdDuplicate
			case "name":
				return productmodel.ErrToppingNameDuplicate
			}
		}
		return common.ErrDB(err)
	}
	return nil
}
