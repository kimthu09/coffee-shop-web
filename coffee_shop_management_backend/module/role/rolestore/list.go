package rolestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) ListRole(
	ctx context.Context) ([]rolemodel.SimpleRole, error) {
	var result []rolemodel.SimpleRole
	db := s.db

	db = db.Table(common.TableRole)

	if err := db.
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
