package cancelnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"context"
)

type ListCancelNoteDetailStore interface {
	ListCancelNoteDetail(
		ctx context.Context,
		cancelNoteId string,
		paging *common.Paging,
	) ([]cancelnotedetailmodel.CancelNoteDetail, error)
}

type listCancelNoteDetailBiz struct {
	store     ListCancelNoteDetailStore
	requester middleware.Requester
}

func NewListCancelNoteDetailBiz(
	store ListCancelNoteDetailStore,
	requester middleware.Requester) *listCancelNoteDetailBiz {
	return &listCancelNoteDetailBiz{store: store, requester: requester}
}

func (biz *listCancelNoteDetailBiz) ListCancelNoteDetail(
	ctx context.Context,
	cancelNoteId string,
	paging *common.Paging) ([]cancelnotedetailmodel.CancelNoteDetail, error) {
	if !biz.requester.IsHasFeature(common.CancelNoteViewFeatureCode) {
		return nil, cancelnotedetailmodel.ErrCancelDetailViewNoPermission
	}

	result, err := biz.store.ListCancelNoteDetail(
		ctx,
		cancelNoteId,
		paging,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
