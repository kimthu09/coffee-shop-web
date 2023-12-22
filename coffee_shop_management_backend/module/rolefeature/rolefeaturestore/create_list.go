package rolefeaturestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
)

func (s *sqlStore) CreateListRoleFeatureDetail(
	ctx context.Context,
	data []rolefeaturemodel.RoleFeature) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
