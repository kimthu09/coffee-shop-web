package exportnoterepo

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
)

type CreateExportNoteStore interface {
	CreateExportNote(
		ctx context.Context,
		data *exportnotemodel.ExportNoteCreate,
	) error
}

type CreateExportNoteDetailStore interface {
	CreateListExportNoteDetail(
		ctx context.Context,
		data []exportnotedetailmodel.ExportNoteDetailCreate,
	) error
}

type UpdateIngredientStore interface {
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientmodel.Ingredient, error)
}

type createExportNoteRepo struct {
	exportNoteStore       CreateExportNoteStore
	exportNoteDetailStore CreateExportNoteDetailStore
	ingredientStore       UpdateIngredientStore
}

func NewCreateExportNoteRepo(
	exportNoteStore CreateExportNoteStore,
	exportNoteDetailStore CreateExportNoteDetailStore,
	ingredientStore UpdateIngredientStore) *createExportNoteRepo {
	return &createExportNoteRepo{
		exportNoteStore:       exportNoteStore,
		exportNoteDetailStore: exportNoteDetailStore,
		ingredientStore:       ingredientStore,
	}
}

func (repo *createExportNoteRepo) HandleExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	if err := repo.exportNoteStore.CreateExportNote(ctx, data); err != nil {
		return err
	}

	if err := repo.exportNoteDetailStore.CreateListExportNoteDetail(
		ctx, data.ExportNoteDetails,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createExportNoteRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	exportNoteId string,
	ingredientTotalAmountNeedUpdate map[string]int) error {
	for key, value := range ingredientTotalAmountNeedUpdate {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": key})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: -value}

		amountLeft := ingredient.Amount - value
		if amountLeft < 0 {
			return exportnotemodel.ErrExportNoteAmountExportIsOverTheStock
		}

		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, key, &ingredientUpdate,
		); err != nil {
			return err
		}
	}
	return nil
}
