package featurestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/feature/featuremodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindFeature(
	ctx context.Context,
	id string) (*featuremodel.Feature, error) {
	var data featuremodel.Feature
	db := s.db

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
