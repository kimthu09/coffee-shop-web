package supplierrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
)

type ListSupplierImportNoteStore interface {
	ListImportNoteBySupplier(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging,
		moreKeys ...string) ([]importnotemodel.ImportNote, error)
}

type seeSupplierImportNoteRepo struct {
	importNoteStore ListSupplierImportNoteStore
}

func NewSeeSupplierImportNoteRepo(
	importNoteStore ListSupplierImportNoteStore) *seeSupplierImportNoteRepo {
	return &seeSupplierImportNoteRepo{
		importNoteStore: importNoteStore,
	}
}

func (biz *seeSupplierImportNoteRepo) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	importNotes, errImportNotes := biz.importNoteStore.ListImportNoteBySupplier(
		ctx,
		supplierId,
		filter,
		paging,
		"CreatedByUser", "ClosedByUser",
	)
	if errImportNotes != nil {
		return nil, errImportNotes
	}

	return importNotes, nil
}
