package importnoterepo

import (
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
)

type CreateImportNoteStore interface {
	CreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
}

type CreateImportNoteDetailStore interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []importnotedetailmodel.ImportNoteDetailCreate,
	) error
}

type UpdatePriceIngredientStore interface {
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientmodel.Ingredient, error)
	UpdatePriceIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdatePrice,
	) error
}

type CheckSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*suppliermodel.Supplier, error)
}

type createImportNoteRepo struct {
	importNoteStore       CreateImportNoteStore
	importNoteDetailStore CreateImportNoteDetailStore
	ingredientStore       UpdatePriceIngredientStore
	supplierStore         CheckSupplierStore
}

func NewCreateImportNoteRepo(
	importNoteStore CreateImportNoteStore,
	importNoteDetailStore CreateImportNoteDetailStore,
	updatePriceIngredientStore UpdatePriceIngredientStore,
	supplierStore CheckSupplierStore) *createImportNoteRepo {
	return &createImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		ingredientStore:       updatePriceIngredientStore,
		supplierStore:         supplierStore,
	}
}

func (repo *createImportNoteRepo) HandleCreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {
	if err := repo.importNoteStore.CreateImportNote(ctx, data); err != nil {
		return err
	}
	if err := repo.importNoteDetailStore.CreateListImportNoteDetail(
		ctx,
		data.ImportNoteDetails); err != nil {
		return err
	}
	return nil
}

func (repo *createImportNoteRepo) UpdatePriceIngredient(
	ctx context.Context,
	ingredientId string,
	price float32) error {
	ingredientUpdatePrice := ingredientmodel.IngredientUpdatePrice{
		Price: &price,
	}

	if err := repo.ingredientStore.UpdatePriceIngredient(
		ctx, ingredientId, &ingredientUpdatePrice,
	); err != nil {
		return err
	}
	return nil
}
