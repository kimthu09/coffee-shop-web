package supplierstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) GetDebtSupplier(
	ctx context.Context,
	supplierId string) (*float32, error) {
	var data suppliermodel.Supplier
	db := s.db

	if err := db.
		Table(common.TableSupplier).
		Select("debt").
		Where("id = ?", supplierId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data.Debt, nil
}
