package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) UpdateStatusFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdateStatus) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
