package invoicestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]invoicemodel.Invoice, error) {
	var result []invoicemodel.Invoice
	db := s.db

	db = db.Table(common.TableInvoice)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("createdAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *invoicemodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.DateFromCreatedAt != nil {
			timeFrom := time.Unix(*filter.DateFromCreatedAt, 0)
			db = db.Where("createdAt >= ?", timeFrom)
		}
		if filter.DateToCreatedAt != nil {
			timeTo := time.Unix(*filter.DateToCreatedAt, 0)
			db = db.Where("createdAt <= ?", timeTo)
		}
		if filter.MinPrice != nil {
			db = db.Where("totalPrice >= ?", filter.MinPrice)
		}
		if filter.MaxPrice != nil {
			db = db.Where("totalPrice <= ?", filter.MaxPrice)
		}
		if filter.CreatedBy != nil {
			db = db.
				Joins("JOIN MUser ON Invoice.createdBy = MUser.id").
				Where("MUser.Id = ?", filter.CreatedBy)
		}
		if filter.Customer != nil {
			db = db.
				Joins("JOIN Customer ON Invoice.customerId = Customer.id").
				Where("Customer.Id = ?", filter.Customer)
		}
	}
}
