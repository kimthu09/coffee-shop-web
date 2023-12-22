package importnoterepo

import (
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
)

type FindImportNoteStore interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*importnotemodel.ImportNote, error)
}

type seeImportNoteDetailRepo struct {
	importNoteStore FindImportNoteStore
}

func NewSeeImportNoteDetailRepo(
	importNoteStore FindImportNoteStore) *seeImportNoteDetailRepo {
	return &seeImportNoteDetailRepo{
		importNoteStore: importNoteStore,
	}
}

func (repo *seeImportNoteDetailRepo) SeeImportNoteDetail(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, errImportNote := repo.importNoteStore.FindImportNote(
		ctx, map[string]interface{}{"id": importNoteId}, "Supplier")
	if errImportNote != nil {
		return nil, errImportNote
	}

	return importNote, nil
}
