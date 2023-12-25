package supplierstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

func (s *sqlStore) UpdateSupplierInfo(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateInfo) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("phone"); key {
			case "phone":
				return suppliermodel.ErrSupplierPhoneDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
