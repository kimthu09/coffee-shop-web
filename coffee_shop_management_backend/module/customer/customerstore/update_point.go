package customerstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateCustomerPoint(
	ctx context.Context,
	id string,
	data *customermodel.CustomerUpdatePoint) error {
	db := s.db

	if err := db.Table(common.TableCustomer).
		Where("id = ?", id).
		Update("point", gorm.Expr("point + ?", data.Amount)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
