package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
)

type ListImportNoteRepo interface {
	ListImportNote(
		ctx context.Context,
		filter *importnotemodel.Filter,
		paging *common.Paging,
	) ([]importnotemodel.ImportNote, error)
}

type listImportNoteBiz struct {
	repo      ListImportNoteRepo
	requester middleware.Requester
}

func NewListImportNoteBiz(
	repo ListImportNoteRepo,
	requester middleware.Requester) *listImportNoteBiz {
	return &listImportNoteBiz{repo: repo, requester: requester}
}

func (biz *listImportNoteBiz) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.ImportNoteViewFeatureCode) {
		return nil, importnotemodel.ErrImportNoteViewNoPermission
	}

	result, err := biz.repo.ListImportNote(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
