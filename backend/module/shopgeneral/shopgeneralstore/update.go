package shopgeneralstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
)

func (s *sqlStore) UpdateGeneralShop(
	ctx context.Context,
	data *shopgeneralmodel.ShopGeneralUpdate) error {
	db := s.db

	if err := db.Where("id = ?", "shop").Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
