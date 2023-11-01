package sizefoodstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
)

func (s *sqlStore) DeleteSizeFood(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db.Begin()

	if err := db.
		Where(conditions).
		Delete(sizefoodmodel.SizeFood{}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
