package userstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateStatusUser(
	ctx context.Context,
	data *usermodel.UserUpdateStatus) error {
	db := s.db

	if err := db.
		Table(common.TableUser).
		Where("id = ?", data.UserId).
		Update("isActive", data.IsActive).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
