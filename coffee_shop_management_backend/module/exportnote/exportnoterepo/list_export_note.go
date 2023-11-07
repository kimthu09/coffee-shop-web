package exportnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
)

type ListExportNoteStore interface {
	ListExportNote(
		ctx context.Context,
		filter *exportnotemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]exportnotemodel.ExportNote, error)
}

type listExportNoteRepo struct {
	store ListExportNoteStore
}

func NewListExportNoteRepo(store ListExportNoteStore) *listExportNoteRepo {
	return &listExportNoteRepo{store: store}
}

func (repo *listExportNoteRepo) ListExportNote(
	ctx context.Context,
	filter *exportnotemodel.Filter,
	paging *common.Paging) ([]exportnotemodel.ExportNote, error) {
	result, err := repo.store.ListExportNote(
		ctx,
		filter,
		[]string{"id", "createBy"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
