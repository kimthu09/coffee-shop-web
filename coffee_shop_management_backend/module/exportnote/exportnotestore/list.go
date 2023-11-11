package exportnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListExportNote(
	ctx context.Context,
	filter *exportnotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]exportnotemodel.ExportNote, error) {
	var result []exportnotemodel.ExportNote
	db := s.db

	db = db.Table(common.TableExportNote)

	handleFilter(db, filter, propertiesContainSearchKey)

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

func handleFilter(
	db *gorm.DB,
	filter *exportnotemodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
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
