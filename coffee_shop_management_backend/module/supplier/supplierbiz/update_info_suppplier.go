package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type UpdateInfoSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
	UpdateSupplierInfo(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateInfo,
	) error
}

type updateInfoSupplierBiz struct {
	store UpdateInfoSupplierStore
}

func NewUpdateInfoSupplierBiz(store UpdateInfoSupplierStore) *updateInfoSupplierBiz {
	return &updateInfoSupplierBiz{store: store}
}

func (biz *updateInfoSupplierBiz) UpdateInfoSupplier(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateInfo) error {

	_, err := biz.store.FindSupplier(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(common.TableCategory, err)
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.UpdateSupplierInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
