package supplierstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateSupplierDebt(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateDebt) error {
	db := s.db

	if err := db.Table(common.TableSupplier).
		Where("id = ?", id).
		Update("debt", gorm.Expr("debt + ?", data.Amount)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
