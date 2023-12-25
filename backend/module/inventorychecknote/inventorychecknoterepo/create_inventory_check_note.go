package inventorychecknoterepo

import (
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
)

type CreateInventoryCheckNoteStore interface {
	CreateInventoryCheckNote(
		ctx context.Context,
		data *inventorychecknotemodel.InventoryCheckNoteCreate,
	) error
}

type CreateInventoryCheckNoteDetailStore interface {
	CreateListInventoryCheckNoteDetail(
		ctx context.Context,
		data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate,
	) error
}

type UpdateIngredientStore interface {
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount) error
	FindIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*ingredientmodel.Ingredient, error)
}

type createInventoryCheckNoteRepo struct {
	inventoryCheckNoteStore       CreateInventoryCheckNoteStore
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore
	ingredientStore               UpdateIngredientStore
}

func NewCreateInventoryCheckNoteRepo(
	inventoryCheckNoteStore CreateInventoryCheckNoteStore,
	inventoryCheckNoteDetailStore CreateInventoryCheckNoteDetailStore,
	ingredientStore UpdateIngredientStore) *createInventoryCheckNoteRepo {
	return &createInventoryCheckNoteRepo{
		inventoryCheckNoteStore:       inventoryCheckNoteStore,
		inventoryCheckNoteDetailStore: inventoryCheckNoteDetailStore,
		ingredientStore:               ingredientStore,
	}
}

func (repo *createInventoryCheckNoteRepo) HandleInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	if err := repo.inventoryCheckNoteStore.CreateInventoryCheckNote(ctx, data); err != nil {
		return err
	}

	if err := repo.inventoryCheckNoteDetailStore.CreateListInventoryCheckNoteDetail(
		ctx, data.Details,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createInventoryCheckNoteRepo) HandleIngredientAmount(
	ctx context.Context,
	data *inventorychecknotemodel.InventoryCheckNoteCreate) error {
	amountDiff := 0
	amountAfter := 0

	for i, value := range data.Details {
		ingredient, errGetIngredient := repo.ingredientStore.FindIngredient(
			ctx, map[string]interface{}{"id": value.IngredientId})
		if errGetIngredient != nil {
			return errGetIngredient
		}

		data.Details[i].Initial = ingredient.Amount
		data.Details[i].Final = ingredient.Amount + value.Difference
		amountDiff += value.Difference
		amountAfter += data.Details[i].Final

		if data.Details[i].Final < 0 {
			return inventorychecknotemodel.ErrInventoryCheckNoteModifyAmountIsInvalid
		}

		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: value.Difference}
		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, value.IngredientId, &ingredientUpdate,
		); err != nil {
			return err
		}
	}

	data.AmountDifferent = amountDiff
	data.AmountAfterAdjust = amountAfter
	return nil
}
