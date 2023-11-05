package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/asyncjob"
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

type payCustomerBiz struct {
	customerStore     PayCustomerStore
	customerDebtStore CreateCustomerDebtStore
	requester         common.Requester
}

func NewUpdatePayBiz(
	customerStore PayCustomerStore,
	customerDebtStore CreateCustomerDebtStore,
	requester common.Requester) *payCustomerBiz {
	return &payCustomerBiz{
		customerStore:     customerStore,
		customerDebtStore: customerDebtStore,
		requester:         requester}
}

func (biz *payCustomerBiz) PayCustomer(
	ctx context.Context,
	idCustomer string,
	data *customermodel.CustomerUpdateDebt) (*string, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	if *data.Amount >= 0 {
		return nil, customermodel.ErrCustomerDebtPayIsInvalid
	}

	debtCurrent, err := biz.customerStore.GetDebtCustomer(
		ctx,
		idCustomer)
	if err != nil {
		return nil, err
	}

	amountPay := *data.Amount
	amountLeft := *debtCurrent + amountPay

	id, errGenerateId := common.GenerateId()
	if errGenerateId != nil {
		return nil, errGenerateId
	}

	debtType := enum.Pay
	customerDebtCreate := customerdebtmodel.CustomerDebtCreate{
		Id:         id,
		IdCustomer: idCustomer,
		Amount:     amountPay,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   biz.requester.GetUserId(),
	}

	jobUpdateCustomer := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.customerStore.UpdateCustomerDebt(ctx, idCustomer, data)
	})
	jobCreateCustomerDebt := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.customerDebtStore.CreateCustomerDebt(ctx, &customerDebtCreate)
	})
	group := asyncjob.NewGroup(
		false,
		jobUpdateCustomer,
		jobCreateCustomerDebt)
	if err := group.Run(context.Background()); err != nil {
		return nil, err
	}

	return &customerDebtCreate.Id, nil
}
