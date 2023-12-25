package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
)

type ListExportNoteRepo interface {
	ListExportNote(
		ctx context.Context,
		filter *exportnotemodel.Filter,
		paging *common.Paging,
	) ([]exportnotemodel.ExportNote, error)
}

type listExportNoteBiz struct {
	repo      ListExportNoteRepo
	requester middleware.Requester
}

func NewListExportNoteRepo(
	repo ListExportNoteRepo,
	requester middleware.Requester) *listExportNoteBiz {
	return &listExportNoteBiz{repo: repo, requester: requester}
}

func (biz *listExportNoteBiz) ListExportNote(
	ctx context.Context,
	filter *exportnotemodel.Filter,
	paging *common.Paging) ([]exportnotemodel.ExportNote, error) {
	if !biz.requester.IsHasFeature(common.ExportNoteViewFeatureCode) {
		return nil, exportnotemodel.ErrExportNoteViewNoPermission
	}

	result, err := biz.repo.ListExportNote(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
