package exportnoterepo

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
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
	GetPriceIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*float32, error)
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
}

type UpdateIngredientDetailStore interface {
	FindIngredientDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientdetailmodel.IngredientDetail, error)
	UpdateIngredientDetail(
		ctx context.Context,
		ingredientId string,
		expiryDate string,
		data *ingredientdetailmodel.IngredientDetailUpdate,
	) error
}

type createExportNoteRepo struct {
	exportNoteStore       CreateExportNoteStore
	exportNoteDetailStore CreateExportNoteDetailStore
	ingredientStore       UpdateIngredientStore
	ingredientDetailStore UpdateIngredientDetailStore
}

func NewCreateExportNoteRepo(
	exportNoteStore CreateExportNoteStore,
	exportNoteDetailStore CreateExportNoteDetailStore,
	ingredientStore UpdateIngredientStore,
	ingredientDetailStore UpdateIngredientDetailStore) *createExportNoteRepo {
	return &createExportNoteRepo{
		exportNoteStore:       exportNoteStore,
		exportNoteDetailStore: exportNoteDetailStore,
		ingredientStore:       ingredientStore,
		ingredientDetailStore: ingredientDetailStore,
	}
}

func (repo *createExportNoteRepo) GetPriceIngredient(
	ctx context.Context,
	ingredientId string) (*float32, error) {
	price, err := repo.ingredientStore.GetPriceIngredient(
		ctx, map[string]interface{}{"id": ingredientId},
	)
	if err != nil {
		return nil, err
	}

	return price, nil
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

func (repo *createExportNoteRepo) HandleIngredientDetail(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	for _, exportNoteDetailCreate := range data.ExportNoteDetails {
		if err := repo.checkIngredientDetail(ctx, &exportNoteDetailCreate); err != nil {
			return err
		}

		if err := repo.updateIngredientDetail(ctx, &exportNoteDetailCreate); err != nil {
			return err
		}
	}
	return nil
}

func (repo *createExportNoteRepo) checkIngredientDetail(
	ctx context.Context,
	data *exportnotedetailmodel.ExportNoteDetailCreate) error {
	currentData, err := repo.ingredientDetailStore.FindIngredientDetail(
		ctx,
		map[string]interface{}{
			"ingredientId": data.IngredientId,
			"expiryDate":   data.ExpiryDate,
		},
	)
	if err != nil {
		return err
	}

	if currentData.Amount < data.AmountExport {
		return exportnotemodel.ErrExportNoteAmountExportIsOverTheStock
	}

	return nil
}

func (repo *createExportNoteRepo) updateIngredientDetail(
	ctx context.Context,
	data *exportnotedetailmodel.ExportNoteDetailCreate) error {
	dataUpdate := ingredientdetailmodel.IngredientDetailUpdate{
		Amount: -data.AmountExport,
	}

	if err := repo.ingredientDetailStore.UpdateIngredientDetail(
		ctx,
		data.IngredientId,
		data.ExpiryDate,
		&dataUpdate,
	); err != nil {
		return err
	}

	return nil
}

func (repo *createExportNoteRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	ingredientTotalAmountNeedUpdate map[string]float32) error {
	for key, value := range ingredientTotalAmountNeedUpdate {
		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: -value}

		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, key, &ingredientUpdate,
		); err != nil {
			return err
		}
	}
	return nil
}
