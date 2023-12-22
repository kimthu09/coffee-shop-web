package importnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"gorm.io/gorm"
	"time"
)

func (s *sqlStore) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging,
	moreKeys ...string) ([]importnotemodel.ImportNote, error) {
	var result []importnotemodel.ImportNote
	db := s.db

	db = db.Table(common.TableImportNote)

	handleSupplierImportFilter(db, filter)

	db = db.Where("supplierId = ?", supplierId)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleSupplierImportFilter(
	db *gorm.DB,
	filterSupplierDebt *filter.SupplierImportFilter) {
	if filterSupplierDebt != nil {
		if filterSupplierDebt.DateFrom != nil {
			timeFrom := time.Unix(*filterSupplierDebt.DateFrom, 0)
			db = db.Where("createdAt >= ?", timeFrom)
		}
		if filterSupplierDebt.DateTo != nil {
			timeTo := time.Unix(*filterSupplierDebt.DateTo, 0)
			db = db.Where("createdAt <= ?", timeTo)
		}
	}
}
