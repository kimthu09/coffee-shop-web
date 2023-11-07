package rolestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
)

func (s *sqlStore) CreateRole(
	ctx context.Context,
	data *rolemodel.RoleCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
