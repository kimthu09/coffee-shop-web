package categorystore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
)

func (s *sqlStore) UpdateInfoCategory(
	ctx context.Context,
	id string,
	data *categorymodel.CategoryUpdateInfo) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("name"); key {
			case "name":
				return categorymodel.ErrCategoryNameDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
