package customerdebtstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListCustomerDebt(
	ctx context.Context,
	customerId string,
	paging *common.Paging) ([]customerdebtmodel.CustomerDebt, error) {
	var result []customerdebtmodel.CustomerDebt
	db := s.db

	db = db.Where("customerId = ?", customerId)

	db = db.Table(common.TableCustomerDebt)

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
