package cancelnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
)

type ListCancelNoteStore interface {
	ListCancelNote(
		ctx context.Context,
		filter *cancelnotemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
	) ([]cancelnotemodel.CancelNote, error)
}

type listCancelNoteRepo struct {
	store ListCancelNoteStore
}

func NewListCancelNoteRepo(store ListCancelNoteStore) *listCancelNoteRepo {
	return &listCancelNoteRepo{store: store}
}

func (repo *listCancelNoteRepo) ListCancelNote(
	ctx context.Context,
	filter *cancelnotemodel.Filter,
	paging *common.Paging) ([]cancelnotemodel.CancelNote, error) {
	result, err := repo.store.ListCancelNote(
		ctx,
		filter,
		[]string{"id", "createBy"},
		paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
