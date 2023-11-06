package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
)

type CreateCustomerStore interface {
	CreateCustomer(
		ctx context.Context,
		data *customermodel.CustomerCreate) error
}

type createCustomerBiz struct {
	store CreateCustomerStore
}

func NewCreateCustomerBiz(store CreateCustomerStore) *createCustomerBiz {
	return &createCustomerBiz{store: store}
}

func (biz *createCustomerBiz) CreateCustomer(
	ctx context.Context,
	data *customermodel.CustomerCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	idAddress, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress

	if err := biz.store.CreateCustomer(ctx, data); err != nil {
		return err
	}

	return nil
}
