package importnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	var result []importnotedetailmodel.ImportNoteDetail
	db := s.db

	db = db.Table(common.TableImportNoteDetail)

	db = db.Where("importNoteId = ?", importNoteId)

	if err := db.
		Preload("Ingredient", func(db *gorm.DB) *gorm.DB {
			return db.Order("Ingredient.name")
		}).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
