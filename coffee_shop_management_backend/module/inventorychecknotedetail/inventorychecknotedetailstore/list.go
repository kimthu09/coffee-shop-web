package inventorychecknotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

func (s *sqlStore) ListInventoryCheckNoteDetail(
	ctx context.Context,
	inventoryCheckNoteId string) ([]inventorychecknotedetailmodel.InventoryCheckNoteDetail, error) {
	var result []inventorychecknotedetailmodel.InventoryCheckNoteDetail
	db := s.db

	db = db.Table(common.TableInventoryCheckNoteDetail)

	db = db.Where("inventoryCheckNoteId = ?", inventoryCheckNoteId)

	if err := db.
		Preload("Ingredient").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
