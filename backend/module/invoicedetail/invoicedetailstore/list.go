package invoicedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListInvoiceDetail(
	ctx context.Context,
	invoiceId string) ([]invoicedetailmodel.InvoiceDetail, error) {
	var result []invoicedetailmodel.InvoiceDetail
	db := s.db

	db = db.Table(common.TableInvoiceDetail)

	db = db.Where("invoiceId = ?", invoiceId)

	if err := db.
		Preload("Food", func(db *gorm.DB) *gorm.DB {
			return db.Order("Food.name desc")
		}).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
