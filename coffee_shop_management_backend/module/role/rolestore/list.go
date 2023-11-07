package rolestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) ListRole(
	ctx context.Context) ([]rolemodel.Role, error) {
	var result []rolemodel.Role
	db := s.db

	db = db.Table(common.TableRole)

	if err := db.
		Preload("RoleFeatures").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
