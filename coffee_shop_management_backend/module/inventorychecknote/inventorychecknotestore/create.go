package inventorychecknotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

func (s *sqlStore) CreateInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
