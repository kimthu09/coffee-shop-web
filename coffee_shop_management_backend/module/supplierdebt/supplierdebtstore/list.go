package supplierdebtstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListSupplierDebt(
	ctx context.Context,
	supplierId string,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	var result []supplierdebtmodel.SupplierDebt
	db := s.db

	db = db.Table(common.TableSupplierDebt)

	db = db.Where("supplierId = ?", supplierId)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("createAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handlePaging(db *gorm.DB, paging *common.Paging) (*gorm.DB, error) {
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset))

	return db, nil
}
