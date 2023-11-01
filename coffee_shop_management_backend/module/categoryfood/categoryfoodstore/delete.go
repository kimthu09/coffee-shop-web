package categoryfoodstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) DeleteCategoryFood(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db.Begin()

	if err := db.Where(conditions).Delete(&categorymodel.Category{}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
