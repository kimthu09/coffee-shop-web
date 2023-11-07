package exportnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"context"
)

type ListExportNoteDetailStore interface {
	ListExportNoteDetail(
		ctx context.Context,
		exportNoteId string,
		paging *common.Paging,
	) ([]exportnotedetailmodel.ExportNoteDetail, error)
}

type listExportNoteDetailBiz struct {
	store     ListExportNoteDetailStore
	requester middleware.Requester
}

func NewListExportNoteDetailBiz(
	store ListExportNoteDetailStore,
	requester middleware.Requester) *listExportNoteDetailBiz {
	return &listExportNoteDetailBiz{store: store, requester: requester}
}

func (biz *listExportNoteDetailBiz) ListExportNoteDetail(
	ctx context.Context,
	exportNoteId string,
	paging *common.Paging) ([]exportnotedetailmodel.ExportNoteDetail, error) {
	if !biz.requester.IsHasFeature(common.ExportNoteViewFeatureCode) {
		return nil, exportnotedetailmodel.ErrExportDetailViewNoPermission
	}

	result, err := biz.store.ListExportNoteDetail(
		ctx,
		exportNoteId,
		paging,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
