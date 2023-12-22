package inventorychecknotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListInventoryCheckNoteDetail(
	ctx context.Context,
	data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
