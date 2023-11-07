package invoicestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"gorm.io/gorm"
	"strings"
)

func (s *sqlStore) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	var result []invoicemodel.Invoice
	db := s.db

	db = db.Table(common.TableInvoice)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("createAt desc").
		Preload("Customer").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func getWhereClause(
	db *gorm.DB,
	searchKey string,
	propertiesContainSearchKey []string) *gorm.DB {
	conditions := make([]string, len(propertiesContainSearchKey))
	args := make([]interface{}, len(propertiesContainSearchKey))

	for i, prop := range propertiesContainSearchKey {
		conditions[i] = prop + " LIKE ?"
		args[i] = "%" + searchKey + "%"
	}

	whereClause := strings.Join(conditions, " OR ")

	return db.Where(whereClause, args...)
}

func handleFilter(
	db *gorm.DB,
	filter *invoicemodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = getWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.IsHasDebt != nil {
			if *filter.IsHasDebt {
				db = db.Where("amountDebt > 0")
			} else {
				db = db.Where("amountDebt = 0")
			}

		}
		if filter.MinPrice != nil {
			db = db.Where("totalPrice >= ?", filter.MinPrice)
		}
		if filter.MaxPrice != nil {
			db = db.Where("totalPrice <= ?", filter.MaxPrice)
		}
	}
}

func handlePaging(db *gorm.DB, paging *common.Paging) (*gorm.DB, error) {
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	db = db.Offset(int(offset))

	return db, nil
}
