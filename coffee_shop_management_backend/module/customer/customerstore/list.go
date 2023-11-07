package customerstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"gorm.io/gorm"
	"strings"
)

func (s *sqlStore) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]customermodel.Customer, error) {
	var result []customermodel.Customer
	db := s.db

	db = db.Table(common.TableCustomer)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("name").
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
	filter *customermodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = getWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.MinDebt != nil {
			db = db.Where("debt >= ?", filter.MinDebt)
		}
		if filter.MaxDebt != nil {
			db = db.Where("debt <= ?", filter.MaxDebt)
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
