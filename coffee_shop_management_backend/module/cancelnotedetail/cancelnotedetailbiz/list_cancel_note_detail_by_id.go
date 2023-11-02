package cancelnotedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"context"
)

type ListCancelNoteDetailByIdStorage interface {
	ListCancelNote(
		ctx context.Context,
		condition map[string]interface{},
		filter *cancelnotedetailmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]cancelnotedetailmodel.CancelNoteDetail, error)
}

type listCancelNoteDetailByIdBiz struct {
	store ListCancelNoteDetailByIdStorage
}

func NewListCancelNoteDetailByIdBiz(store ListCancelNoteDetailByIdStorage) *listCancelNoteDetailByIdBiz {
	return &listCancelNoteDetailByIdBiz{store: store}
}

func (biz *listCancelNoteDetailByIdBiz) ListCancelNoteDetailByIdBiz(
	ctx context.Context,
	cancelNoteId string,
	filter *cancelnotedetailmodel.Filter,
	paging *common.Paging) ([]cancelnotedetailmodel.CancelNoteDetail, error) {
	result, err := biz.store.ListCancelNote(
		ctx, map[string]interface{}{"cancelNoteId": cancelNoteId}, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(common.TableCategory, err)
	}

	return result, nil
}
