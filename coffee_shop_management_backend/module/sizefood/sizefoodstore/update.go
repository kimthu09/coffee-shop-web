package sizefoodstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

func (s *sqlStore) UpdateSizeFood(
	ctx context.Context,
	foodId string,
	sizeId string,
	data *sizefoodmodel.SizeFoodUpdate) error {
	db := s.db.Begin()

	if err := db.
		Where("foodId = ? and sizeId = ?", foodId, sizeId).
		Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
