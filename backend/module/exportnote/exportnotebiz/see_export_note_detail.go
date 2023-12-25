package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
)

type SeeExportNoteDetailRepo interface {
	SeeExportNoteDetail(
		ctx context.Context,
		exportNoteId string,
	) (*exportnotemodel.ExportNote, error)
}

type seeExportNoteDetailBiz struct {
	repo      SeeExportNoteDetailRepo
	requester middleware.Requester
}

func NewSeeExportNoteDetailBiz(
	repo SeeExportNoteDetailRepo,
	requester middleware.Requester) *seeExportNoteDetailBiz {
	return &seeExportNoteDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeExportNoteDetailBiz) SeeExportNoteDetail(
	ctx context.Context,
	exportNoteId string) (*exportnotemodel.ExportNote, error) {
	if !biz.requester.IsHasFeature(common.ExportNoteViewFeatureCode) {
		return nil, exportnotemodel.ErrExportNoteViewNoPermission
	}

	exportNote, errExportNote := biz.repo.SeeExportNoteDetail(
		ctx, exportNoteId)
	if errExportNote != nil {
		return nil, errExportNote
	}

	return exportNote, nil
}
