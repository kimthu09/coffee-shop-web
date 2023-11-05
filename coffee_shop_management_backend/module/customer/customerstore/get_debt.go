package customerstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) GetDebtCustomer(
	ctx context.Context,
	customerId string) (*float32, error) {
	var data customermodel.Customer
	db := s.db

	if err := db.
		Table(common.TableCustomer).
		Select("debt").
		Where("id = ?", customerId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data.Debt, nil
}
