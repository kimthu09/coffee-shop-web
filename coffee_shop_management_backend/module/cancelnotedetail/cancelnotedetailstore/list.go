package cancelnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) ListCancelNoteDetail(
	ctx context.Context,
	cancelNoteId string,
	paging *common.Paging) ([]cancelnotedetailmodel.CancelNoteDetail, error) {
	var result []cancelnotedetailmodel.CancelNoteDetail
	db := s.db

	db = db.Table(common.TableCancelNoteDetail)

	db = db.Where("cancelNoteId = ?", cancelNoteId)

	dbTemp, errPaging := handlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Preload("Ingredient", func(db *gorm.DB) *gorm.DB {
			return db.Order("Ingredient.name")
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
