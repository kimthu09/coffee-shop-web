package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
)

type SeeImportNoteDetailRepo interface {
	SeeImportNoteDetail(
		ctx context.Context,
		importNoteId string,
	) (*importnotemodel.ImportNote, error)
}

type seeImportNoteDetailBiz struct {
	repo      SeeImportNoteDetailRepo
	requester middleware.Requester
}

func NewSeeImportNoteDetailBiz(
	repo SeeImportNoteDetailRepo,
	requester middleware.Requester) *seeImportNoteDetailBiz {
	return &seeImportNoteDetailBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeImportNoteDetailBiz) SeeImportNoteDetail(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	if !biz.requester.IsHasFeature(common.ImportNoteViewFeatureCode) {
		return nil, importnotemodel.ErrImportNoteViewNoPermission
	}

	importNote, errImportNote := biz.repo.SeeImportNoteDetail(
		ctx, importNoteId)
	if errImportNote != nil {
		return nil, errImportNote
	}

	return importNote, nil
}
