package customerdebtstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"context"
)

func (s *sqlStore) CreateCustomerDebt(
	ctx context.Context,
	data *customerdebtmodel.CustomerDebtCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
