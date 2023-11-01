package cancelnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
)

func (s *sqlStore) CreateCancelNote(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate) error {
	db := s.db.Begin()

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
