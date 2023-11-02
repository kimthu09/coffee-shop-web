package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindIngredient(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*ingredientmodel.Ingredient, error) {
	var data ingredientmodel.Ingredient
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
