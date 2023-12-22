package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListImportNoteBySupplierRepo interface {
	ListImportNoteBySupplier(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging) ([]importnotemodel.ImportNote, error)
}

type listImportNoteBySupplierBiz struct {
	repo      ListImportNoteBySupplierRepo
	requester middleware.Requester
}

func NewListImportNoteBySupplierBiz(
	repo ListImportNoteBySupplierRepo,
	requester middleware.Requester) *listImportNoteBySupplierBiz {
	return &listImportNoteBySupplierBiz{repo: repo, requester: requester}
}

func (biz *listImportNoteBySupplierBiz) ListImportNoteBySupplier(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	result, err := biz.repo.ListImportNoteBySupplier(ctx, supplierId, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
