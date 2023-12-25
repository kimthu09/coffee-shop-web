package importnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

func (s *sqlStore) CreateListImportNoteDetail(
	ctx context.Context,
	data []importnotedetailmodel.ImportNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
