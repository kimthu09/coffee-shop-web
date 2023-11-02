package supplierdebtstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

func (s *sqlStore) CreateSupplierDebt(
	ctx context.Context,
	data *supplierdebtmodel.SupplierDebtCreate) error {
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
