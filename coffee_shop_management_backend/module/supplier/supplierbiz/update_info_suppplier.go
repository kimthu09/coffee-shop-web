package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type UpdateInfoSupplierRepo interface {
	CheckExist(
		ctx context.Context,
		supplierId string,
	) error
	UpdateSupplierInfo(
		ctx context.Context,
		supplierId string,
		data *suppliermodel.SupplierUpdateInfo,
	) error
}

type updateInfoSupplierBiz struct {
	repo      UpdateInfoSupplierRepo
	requester middleware.Requester
}

func NewUpdateInfoSupplierBiz(
	repo UpdateInfoSupplierRepo,
	requester middleware.Requester) *updateInfoSupplierBiz {
	return &updateInfoSupplierBiz{repo: repo, requester: requester}
}

func (biz *updateInfoSupplierBiz) UpdateInfoSupplier(
	ctx context.Context,
	id string,
	data *suppliermodel.SupplierUpdateInfo) error {
	if !biz.requester.IsHasFeature(common.SupplierUpdateInfoFeatureCode) {
		return suppliermodel.ErrSupplierUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckExist(ctx, id); err != nil {
		return err
	}

	if err := biz.repo.UpdateSupplierInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
