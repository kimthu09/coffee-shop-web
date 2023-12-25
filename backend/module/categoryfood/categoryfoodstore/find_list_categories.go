package categoryfoodstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindListCategories(
	ctx context.Context,
	foodId string) ([]categorymodel.SimpleCategoryWithId, error) {
	var data []categorymodel.SimpleCategoryWithId
	db := s.db

	if err := db.
		Table(common.TableCategoryFood).
		Where("foodId = ?", foodId).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
