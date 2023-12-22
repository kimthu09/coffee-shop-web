package supplierbiz

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type GetAllSupplierRepo interface {
	GetAllSupplier(
		ctx context.Context) ([]suppliermodel.SimpleSupplier, error)
}

type getAllSupplierBiz struct {
	repo GetAllSupplierRepo
}

func NewGetAllSupplierBiz(
	repo GetAllSupplierRepo) *getAllSupplierBiz {
	return &getAllSupplierBiz{repo: repo}
}

func (biz *getAllSupplierBiz) GetAllUser(
	ctx context.Context) ([]suppliermodel.SimpleSupplier, error) {
	result, err := biz.repo.GetAllSupplier(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
