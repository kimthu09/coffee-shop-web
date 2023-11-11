package importnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

type ListImportNoteDetailStore interface {
	ListImportNoteDetail(
		ctx context.Context,
		importNoteId string,
		paging *common.Paging,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type listImportNoteDetailBiz struct {
	store     ListImportNoteDetailStore
	requester middleware.Requester
}

func NewListImportNoteDetailBiz(
	store ListImportNoteDetailStore,
	requester middleware.Requester) *listImportNoteDetailBiz {
	return &listImportNoteDetailBiz{store: store, requester: requester}
}

func (biz *listImportNoteDetailBiz) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) ([]importnotedetailmodel.ImportNoteDetail, error) {
	if !biz.requester.IsHasFeature(common.ImportNoteViewFeatureCode) {
		return nil, importnotedetailmodel.ErrImportDetailViewNoPermission
	}

	result, err := biz.store.ListImportNoteDetail(
		ctx,
		importNoteId,
		paging,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
