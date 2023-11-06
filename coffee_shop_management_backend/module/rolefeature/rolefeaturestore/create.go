package rolefeaturestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
)

func (s *sqlStore) CreateRoleFeature(
	ctx context.Context,
	data *rolefeaturemodel.RoleFeature) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
