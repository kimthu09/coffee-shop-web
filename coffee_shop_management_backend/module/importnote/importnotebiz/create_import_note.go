package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"context"
)

type CreateImportNoteRepo interface {
	CheckIngredient(
		ctx context.Context,
		ingredientId string,
	) error
	HandleCreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
	UpdatePriceIngredient(
		ctx context.Context,
		ingredientId string,
		price float32,
	) error
	CheckSupplier(
		ctx context.Context,
		supplierId string,
	) error
}

type createImportNoteBiz struct {
	repo CreateImportNoteRepo
}

func NewCreateImportNoteBiz(
	repo CreateImportNoteRepo) *createImportNoteBiz {
	return &createImportNoteBiz{
		repo: repo,
	}
}

func (biz *createImportNoteBiz) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	for _, v := range data.ImportNoteDetails {
		if err := biz.repo.CheckIngredient(ctx, v.IngredientId); err != nil {
			return err
		}
	}

	if err := handleImportNoteCreateId(data); err != nil {
		return err
	}

	if err := biz.repo.CheckSupplier(ctx, data.SupplierId); err != nil {
		return err
	}

	handleTotalPrice(data)

	if err := biz.repo.HandleCreateImportNote(ctx, data); err != nil {
		return err
	}

	for _, v := range data.ImportNoteDetails {
		if v.IsReplacePrice {
			if err := biz.repo.UpdatePriceIngredient(
				ctx, v.IngredientId, v.Price,
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func handleImportNoteCreateId(data *importnotemodel.ImportNoteCreate) error {
	idImportNote, err := common.IdProcess(data.Id)
	if err != nil {
		return err
	}
	data.Id = idImportNote

	for i := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].ImportNoteId = *idImportNote
	}
	return nil
}

func handleTotalPrice(data *importnotemodel.ImportNoteCreate) {
	var totalPrice float32 = 0
	for _, importNoteDetail := range data.ImportNoteDetails {
		totalPrice += importNoteDetail.Price * importNoteDetail.AmountImport
	}
	data.TotalPrice = totalPrice
}
