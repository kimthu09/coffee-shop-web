package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
)

func (s *sqlStore) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdate,
) error {

	db := s.db.Begin()

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
