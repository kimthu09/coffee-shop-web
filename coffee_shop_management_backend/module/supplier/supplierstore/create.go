package supplierstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

func (s *sqlStore) CreateSupplier(ctx context.Context, data *suppliermodel.SupplierCreate) error {
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
