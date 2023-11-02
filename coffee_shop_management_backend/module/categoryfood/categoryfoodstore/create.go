package categoryfoodstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
	"context"
)

func (s *sqlStore) CreateCategoryFood(
	ctx context.Context,
	data *categoryfoodmodel.CategoryFoodCreate) error {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
