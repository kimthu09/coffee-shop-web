package categoryfoodstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
	"context"
)

func (s *sqlStore) CreateCategoryFood(
	ctx context.Context,
	data *categoryfoodmodel.CategoryFoodCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
