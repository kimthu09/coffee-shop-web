package customerstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindCustomer(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*customermodel.Customer, error) {
	var data customermodel.Customer
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
