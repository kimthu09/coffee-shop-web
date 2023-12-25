package exportnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListExportNoteDetail(
	ctx context.Context,
	exportNoteId string) ([]exportnotedetailmodel.ExportNoteDetail, error) {
	var result []exportnotedetailmodel.ExportNoteDetail
	db := s.db

	db = db.Table(common.TableExportNoteDetail)

	db = db.Where("exportNoteId = ?", exportNoteId)

	if err := db.
		Preload("Ingredient", func(db *gorm.DB) *gorm.DB {
			return db.Order("Ingredient.name")
		}).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
