package userstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateInfoUser(
	ctx context.Context,
	id string,
	data *usermodel.UserUpdateInfo) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
