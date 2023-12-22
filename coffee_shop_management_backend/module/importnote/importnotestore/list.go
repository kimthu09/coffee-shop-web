package importnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string) ([]importnotemodel.ImportNote, error) {
	var result []importnotemodel.ImportNote
	db := s.db

	db = db.Table(common.TableImportNote)

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
	filter *importnotemodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
		if filter.Status != "" {
			db = db.Where("status = ?", filter.Status)
		}
		if filter.MinPrice != nil {
			db = db.Where("totalPrice >= ?", filter.MinPrice)
		}
		if filter.MaxPrice != nil {
			db = db.Where("totalPrice <= ?", filter.MaxPrice)
		}
	}
}
