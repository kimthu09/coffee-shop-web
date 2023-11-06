package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type CreateSupplierStore interface {
	CreateSupplier(ctx context.Context, data *suppliermodel.SupplierCreate) error
}

type createSupplierBiz struct {
	store CreateSupplierStore
}

func NewCreateSupplierBiz(store CreateSupplierStore) *createSupplierBiz {
	return &createSupplierBiz{store: store}
}

func (biz *createSupplierBiz) CreateSupplier(
	ctx context.Context,
	data *suppliermodel.SupplierCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	idAddress, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = idAddress

	if err := biz.store.CreateSupplier(ctx, data); err != nil {
		return err
	}

	return nil
}
