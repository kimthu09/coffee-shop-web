package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) UpdateFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdateInfo) error {

	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return productmodel.ErrFoodNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
