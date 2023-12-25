package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
)

type SeeSupplierImportNoteRepo interface {
	SeeSupplierImportNote(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging) ([]importnotemodel.ImportNote, error)
}

type seeSupplierImportNoteBiz struct {
	repo      SeeSupplierImportNoteRepo
	requester middleware.Requester
}

func NewSeeSupplierImportNoteBiz(
	repo SeeSupplierImportNoteRepo,
	requester middleware.Requester) *seeSupplierImportNoteBiz {
	return &seeSupplierImportNoteBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierImportNoteBiz) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	importNotes, err := biz.repo.SeeSupplierImportNote(
		ctx, supplierId, filter, paging)
	if err != nil {
		return nil, err
	}

	return importNotes, nil
}
