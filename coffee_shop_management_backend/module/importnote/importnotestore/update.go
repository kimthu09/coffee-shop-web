package importnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
)

func (s *sqlStore) UpdateImportNote(
	ctx context.Context,
	id string,
	data *importnotemodel.ImportNoteUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
