package invoicedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListInvoiceDetail(
	ctx context.Context,
	invoiceId string,
	paging *common.Paging) ([]invoicedetailmodel.InvoiceDetail, error) {
	var result []invoicedetailmodel.InvoiceDetail
	db := s.db

	db = db.Table(common.TableInvoiceDetail)

	db = db.Where("invoiceId = ?", invoiceId)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Preload("Food", func(db *gorm.DB) *gorm.DB {
			return db.Order("Food.name desc")
		}).
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
