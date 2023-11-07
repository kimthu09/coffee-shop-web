package cancelnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
)

type ListCancelNoteRepo interface {
	ListCancelNote(
		ctx context.Context,
		filter *cancelnotemodel.Filter,
		paging *common.Paging,
	) ([]cancelnotemodel.CancelNote, error)
}

type listCancelNoteBiz struct {
	requester middleware.Requester
	repo      ListCancelNoteRepo
}

func NewListCancelNoteRepo(
	repo ListCancelNoteRepo,
	requester middleware.Requester) *listCancelNoteBiz {
	return &listCancelNoteBiz{repo: repo, requester: requester}
}

func (biz *listCancelNoteBiz) ListCancelNote(
	ctx context.Context,
	filter *cancelnotemodel.Filter,
	paging *common.Paging) ([]cancelnotemodel.CancelNote, error) {
	if !biz.requester.IsHasFeature(common.CancelNoteViewFeatureCode) {
		return nil, cancelnotemodel.ErrCancelNoteViewNoPermission
	}

	result, err := biz.repo.ListCancelNote(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
