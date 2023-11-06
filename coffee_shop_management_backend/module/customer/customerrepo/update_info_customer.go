package customerrepo

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
)

type UpdateInfoCustomerStore interface {
	FindCustomer(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
	UpdateCustomerInfo(
		ctx context.Context,
		id string,
		data *customermodel.CustomerUpdateInfo,
	) error
}

type updateInfoCustomerRepo struct {
	store UpdateInfoCustomerStore
}

func NewUpdateInfoSupplierRepo(store UpdateInfoCustomerStore) *updateInfoCustomerRepo {
	return &updateInfoCustomerRepo{store: store}
}

func (repo *updateInfoCustomerRepo) CheckExist(
	ctx context.Context,
	customerId string) error {
	if _, err := repo.store.FindCustomer(
		ctx,
		map[string]interface{}{
			"id": customerId,
		},
	); err != nil {
		return err
	}
	return nil
}

func (repo *updateInfoCustomerRepo) UpdateCustomerInfo(
	ctx context.Context,
	customerId string,
	data *customermodel.CustomerUpdateInfo) error {
	if err := repo.store.UpdateCustomerInfo(ctx, customerId, data); err != nil {
		return err
	}
	return nil
}
