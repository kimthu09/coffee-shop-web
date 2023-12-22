package exportnoterepo

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"context"
)

type SeeExportNoteDetailStore interface {
	ListExportNoteDetail(
		ctx context.Context,
		exportNoteId string) ([]exportnotedetailmodel.ExportNoteDetail, error)
}

type FindExportNoteStore interface {
	FindExportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*exportnotemodel.ExportNote, error)
}

type seeExportNoteDetailRepo struct {
	exportNoteDetailStore SeeExportNoteDetailStore
	exportNoteStore       FindExportNoteStore
}

func NewSeeExportNoteDetailRepo(
	exportNoteDetailStore SeeExportNoteDetailStore,
	exportNoteStore FindExportNoteStore) *seeExportNoteDetailRepo {
	return &seeExportNoteDetailRepo{
		exportNoteDetailStore: exportNoteDetailStore,
		exportNoteStore:       exportNoteStore,
	}
}

func (biz *seeExportNoteDetailRepo) SeeExportNoteDetail(
	ctx context.Context,
	exportNoteId string) (*exportnotemodel.ExportNote, error) {
	exportNote, errExportNote := biz.exportNoteStore.FindExportNote(
		ctx, map[string]interface{}{"id": exportNoteId}, "CreatedByUser")
	if errExportNote != nil {
		return nil, errExportNote
	}

	details, errExportNoteDetail := biz.exportNoteDetailStore.ListExportNoteDetail(
		ctx,
		exportNoteId,
	)
	if errExportNoteDetail != nil {
		return nil, errExportNoteDetail
	}

	exportNote.Details = details

	return exportNote, nil
}
