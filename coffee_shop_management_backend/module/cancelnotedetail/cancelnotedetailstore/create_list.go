package cancelnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListCancelNoteDetail(
	ctx context.Context,
	data []cancelnotedetailmodel.CancelNoteDetailCreate) error {
	db := s.db.Begin()

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
