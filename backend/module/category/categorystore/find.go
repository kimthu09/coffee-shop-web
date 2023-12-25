package categorystore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindCategory(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*categorymodel.Category, error) {
	var data categorymodel.Category
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
