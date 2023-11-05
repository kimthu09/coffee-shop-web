package cancelnoterepo

import (
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
)

type CreateCancelNoteStore interface {
	CreateCancelNote(
		ctx context.Context,
		data *cancelnotemodel.CancelNoteCreate,
	) error
}

type CreateCancelNoteDetailStore interface {
	CreateListCancelNoteDetail(
		ctx context.Context,
		data []cancelnotedetailmodel.CancelNoteDetailCreate,
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

type createCancelNoteRepo struct {
	cancelNoteStore       CreateCancelNoteStore
	cancelNoteDetailStore CreateCancelNoteDetailStore
	ingredientStore       UpdateIngredientStore
	ingredientDetailStore UpdateIngredientDetailStore
}

func NewCreateCancelNoteRepo(
	cancelNoteStore CreateCancelNoteStore,
	cancelNoteDetailStore CreateCancelNoteDetailStore,
	ingredientStore UpdateIngredientStore,
	ingredientDetailStore UpdateIngredientDetailStore) *createCancelNoteRepo {
	return &createCancelNoteRepo{
		cancelNoteStore:       cancelNoteStore,
		cancelNoteDetailStore: cancelNoteDetailStore,
		ingredientStore:       ingredientStore,
		ingredientDetailStore: ingredientDetailStore,
	}
}

func (repo *createCancelNoteRepo) GetPriceIngredient(
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

func (repo *createCancelNoteRepo) HandleCancelNote(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate) error {
	if err := repo.cancelNoteStore.CreateCancelNote(ctx, data); err != nil {
		return err
	}

	if err := repo.cancelNoteDetailStore.CreateListCancelNoteDetail(
		ctx, data.CancelNoteCreateDetails,
	); err != nil {
		return err
	}
	return nil
}

func (repo *createCancelNoteRepo) HandleIngredientDetail(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate) error {
	for _, cancelNoteDetailCreate := range data.CancelNoteCreateDetails {
		if err := repo.checkIngredientDetail(ctx, &cancelNoteDetailCreate); err != nil {
			return err
		}

		if err := repo.updateIngredientDetail(ctx, &cancelNoteDetailCreate); err != nil {
			return err
		}
	}
	return nil
}

func (repo *createCancelNoteRepo) checkIngredientDetail(
	ctx context.Context,
	data *cancelnotedetailmodel.CancelNoteDetailCreate) error {
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

	if currentData.Amount < data.AmountCancel {
		return cancelnotemodel.ErrCancelNoteAmountCancelIsOverTheStock
	}

	return nil
}

func (repo *createCancelNoteRepo) updateIngredientDetail(
	ctx context.Context,
	data *cancelnotedetailmodel.CancelNoteDetailCreate) error {
	dataUpdate := ingredientdetailmodel.IngredientDetailUpdate{
		Amount: -data.AmountCancel,
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

func (repo *createCancelNoteRepo) HandleIngredientTotalAmount(
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
