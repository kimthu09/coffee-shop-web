package inventorychecknotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
)

type CreateInventoryCheckNoteRepo interface {
	HandleInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.InventoryCheckNoteCreate) error
	HandleIngredientAmount(
		ctx context.Context,
		data *inventorychecknotemodel.InventoryCheckNoteCreate) error
}

type createInventoryCheckNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateInventoryCheckNoteRepo
	requester middleware.Requester
}

func NewCreateInventoryCheckNoteBiz(
	gen generator.IdGenerator,
	repo CreateInventoryCheckNoteRepo,
	requester middleware.Requester) *createInventoryCheckNoteBiz {
	return &createInventoryCheckNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createInventoryCheckNoteBiz) CreateInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	if !biz.requester.IsHasFeature(common.InventoryCheckNoteCreateFeatureCode) {
		return inventorychecknotemodel.ErrInventoryCheckNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInventoryCheckNoteId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleIngredientAmount(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleInventoryCheckNote(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleInventoryCheckNoteId(
	gen generator.IdGenerator,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	id, errGenerateId := gen.IdProcess(data.Id)
	if errGenerateId != nil {
		return errGenerateId
	}
	data.Id = id

	for i := range data.Details {
		data.Details[i].InventoryCheckNoteId = *id
	}

	return nil
}
