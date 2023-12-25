package importnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
)

type ListImportNoteStore interface {
	ListImportNote(
		ctx context.Context,
		filter *importnotemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string) ([]importnotemodel.ImportNote, error)
}

type listImportNoteRepo struct {
	store ListImportNoteStore
}

func NewListImportNoteRepo(store ListImportNoteStore) *listImportNoteRepo {
	return &listImportNoteRepo{store: store}
}

func (repo *listImportNoteRepo) ListImportNote(
	ctx context.Context,
	filter *importnotemodel.Filter,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	result, err := repo.store.ListImportNote(
		ctx,
		filter,
		[]string{"ImportNote.id"},
		paging,
		"Supplier", "CreatedByUser", "ClosedByUser")

	if err != nil {
		return nil, err
	}

	return result, nil
}
