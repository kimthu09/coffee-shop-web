package customerbiz

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

type updateInfoCustomerBiz struct {
	store UpdateInfoCustomerStore
}

func NewUpdateInfoSupplierBiz(store UpdateInfoCustomerStore) *updateInfoCustomerBiz {
	return &updateInfoCustomerBiz{store: store}
}

func (biz *updateInfoCustomerBiz) UpdateInfoCustomer(
	ctx context.Context,
	id string,
	data *customermodel.CustomerUpdateInfo) error {

	if err := data.Validate(); err != nil {
		return err
	}

	_, err := biz.store.FindCustomer(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.UpdateCustomerInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
