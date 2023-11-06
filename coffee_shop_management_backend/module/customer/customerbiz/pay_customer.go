package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/customerdebt/customerdebtmodel"
	"context"
)

type PayCustomerRepo interface {
	GetDebtCustomer(
		ctx context.Context,
		customerId string,
	) (*float32, error)
	CreateCustomerDebt(
		ctx context.Context,
		data *customerdebtmodel.CustomerDebtCreate,
	) error
	UpdateDebtCustomer(
		ctx context.Context,
		customerId string,
		data *customermodel.CustomerUpdateDebt,
	) error
}

type payCustomerBiz struct {
	gen       generator.IdGenerator
	repo      PayCustomerRepo
	requester middleware.Requester
}

func NewUpdatePayBiz(
	gen generator.IdGenerator,
	repo PayCustomerRepo,
	requester middleware.Requester) *payCustomerBiz {
	return &payCustomerBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *payCustomerBiz) PayCustomer(
	ctx context.Context,
	customerId string,
	data *customermodel.CustomerUpdateDebt) (*string, error) {
	if !biz.requester.IsHasFeature(common.CustomerPayFeatureCode) {
		return nil, customermodel.ErrCustomerPayNoPermission
	}

	if err := validateCustomerUpdateDebt(data); err != nil {
		return nil, err
	}

	debtCurrent, errGetDebt := biz.repo.GetDebtCustomer(ctx, customerId)
	if errGetDebt != nil {
		return nil, errGetDebt
	}

	customerDebtCreate, errGetCustomerDebtCreate := getCustomerDebtCreate(
		biz.gen, customerId, *debtCurrent, data,
	)
	if errGetCustomerDebtCreate != nil {
		return nil, errGetCustomerDebtCreate
	}

	if err := biz.repo.UpdateDebtCustomer(ctx, customerId, data); err != nil {
		return nil, err
	}

	if err := biz.repo.CreateCustomerDebt(ctx, customerDebtCreate); err != nil {
		return nil, err
	}

	return &customerDebtCreate.Id, nil
}

func validateCustomerUpdateDebt(data *customermodel.CustomerUpdateDebt) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if *data.Amount >= 0 {
		return customermodel.ErrCustomerDebtPayIsInvalid
	}

	return nil
}

func getCustomerDebtCreate(
	gen generator.IdGenerator,
	customerId string,
	currentDebt float32,
	data *customermodel.CustomerUpdateDebt,
) (*customerdebtmodel.CustomerDebtCreate, error) {
	amountPay := *data.Amount
	amountLeft := currentDebt + amountPay

	id, err := gen.GenerateId()
	if err != nil {
		return nil, err
	}

	debtType := enum.Pay
	customerDebtCreate := customerdebtmodel.CustomerDebtCreate{
		Id:         id,
		IdCustomer: customerId,
		Amount:     amountPay,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   data.CreateBy,
	}

	return &customerDebtCreate, nil
}
