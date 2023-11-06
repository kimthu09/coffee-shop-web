package customerrepo

import (
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"context"
)

type PayCustomerStore interface {
	GetDebtCustomer(
		ctx context.Context,
		customerId string,
	) (*float32, error)
	UpdateCustomerDebt(
		ctx context.Context,
		id string,
		data *customermodel.CustomerUpdateDebt,
	) error
}

type CreateCustomerDebtStore interface {
	CreateCustomerDebt(
		ctx context.Context,
		data *customerdebtmodel.CustomerDebtCreate,
	) error
}

type payCustomerRepo struct {
	customerStore     PayCustomerStore
	customerDebtStore CreateCustomerDebtStore
}

func NewUpdatePayRepo(
	customerStore PayCustomerStore,
	customerDebtStore CreateCustomerDebtStore) *payCustomerRepo {
	return &payCustomerRepo{
		customerStore:     customerStore,
		customerDebtStore: customerDebtStore}
}

func (repo *payCustomerRepo) GetDebtCustomer(
	ctx context.Context,
	customerId string) (*float32, error) {
	debtCurrent, err := repo.customerStore.GetDebtCustomer(
		ctx,
		customerId)
	if err != nil {
		return nil, err
	}

	return debtCurrent, nil
}

func (repo *payCustomerRepo) CreateCustomerDebt(
	ctx context.Context,
	data *customerdebtmodel.CustomerDebtCreate) error {
	if err := repo.customerDebtStore.CreateCustomerDebt(ctx, data); err != nil {
		return err
	}
	return nil
}

func (repo *payCustomerRepo) UpdateDebtCustomer(
	ctx context.Context,
	customerId string,
	data *customermodel.CustomerUpdateDebt) error {
	if err := repo.customerStore.UpdateCustomerDebt(ctx, customerId, data); err != nil {
		return err
	}
	return nil
}
