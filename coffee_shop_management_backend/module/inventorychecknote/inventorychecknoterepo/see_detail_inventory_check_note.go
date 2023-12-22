package inventorychecknoterepo

import (
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

type SeeDetailInventoryCheckNoteStore interface {
	ListInventoryCheckNoteDetail(
		ctx context.Context,
		inventoryCheckNoteId string) ([]inventorychecknotedetailmodel.InventoryCheckNoteDetail, error)
}

type FindInventoryCheckNoteStore interface {
	FindInventoryCheckNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*inventorychecknotemodel.InventoryCheckNote, error)
}

type seeDetailInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       FindInventoryCheckNoteStore
	inventoryCheckNoteDetailStore SeeDetailInventoryCheckNoteStore
}

func NewSeeDetailInventoryCheckNoteRepo(
	inventoryCheckNoteStore FindInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore SeeDetailInventoryCheckNoteStore) *seeDetailInventoryCheckNoteRepo {
	return &seeDetailInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
	}
}

func (repo *seeDetailInventoryCheckNoteRepo) SeeDetailInventoryCheckNote(
	ctx context.Context,
	inventoryCheckNoteId string) (*inventorychecknotemodel.InventoryCheckNote, error) {
	inventoryCheckNote, errInventoryCheckNote :=
		repo.inventoryCheckNoteStore.FindInventoryCheckNote(
			ctx,
			map[string]interface{}{"id": inventoryCheckNoteId},
			"CreatedByUser")
	if errInventoryCheckNote != nil {
		return nil, errInventoryCheckNote
	}

	details, errInventoryCheckNoteDetail := repo.inventoryCheckNoteDetailStore.ListInventoryCheckNoteDetail(
		ctx,
		inventoryCheckNoteId,
	)
	if errInventoryCheckNoteDetail != nil {
		return nil, errInventoryCheckNoteDetail
	}

	inventoryCheckNote.Details = details

	return inventoryCheckNote, nil
}
