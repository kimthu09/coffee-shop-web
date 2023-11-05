package cancelnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListCancelNoteDetail(
	ctx context.Context,
	data []cancelnotedetailmodel.CancelNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
