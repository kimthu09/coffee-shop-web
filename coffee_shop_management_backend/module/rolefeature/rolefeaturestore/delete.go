package rolefeaturestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) DeleteRoleFeature(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.
		Table(common.TableRoleFeature).
		Where(conditions).
		Delete(&categorymodel.Category{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
