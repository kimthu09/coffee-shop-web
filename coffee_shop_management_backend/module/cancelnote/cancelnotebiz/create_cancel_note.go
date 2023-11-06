package cancelnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"context"
)

type CreateCancelNoteRepo interface {
	GetPriceIngredient(
		ctx context.Context,
		ingredientId string,
	) (*float32, error)
	HandleCancelNote(
		ctx context.Context,
		data *cancelnotemodel.CancelNoteCreate,
	) error
	HandleIngredientDetail(
		ctx context.Context,
		data *cancelnotemodel.CancelNoteCreate,
	) error
	HandleIngredientTotalAmount(
		ctx context.Context,
		ingredientTotalAmountNeedUpdate map[string]float32,
	) error
}

type createCancelNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateCancelNoteRepo
	requester middleware.Requester
}

func NewCreateCancelNoteBiz(
	generator generator.IdGenerator,
	repo CreateCancelNoteRepo,
	requester middleware.Requester) *createCancelNoteBiz {
	return &createCancelNoteBiz{
		gen:       generator,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createCancelNoteBiz) CreateCancelNote(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate) error {
	if !biz.requester.IsHasFeature(common.CancelNoteCreateFeatureCode) {
		return cancelnotemodel.ErrCancelNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleCancelNoteId(biz.gen, data); err != nil {
		return err
	}

	mapIngredient := getMapIngredientExist(data)
	var totalPrice float32 = 0
	for ingredientId, totalAmountOfIngredientId := range mapIngredient {
		price, err := biz.repo.GetPriceIngredient(
			ctx, ingredientId,
		)
		if err != nil {
			return err
		}

		totalPrice += *price * totalAmountOfIngredientId
	}
	data.TotalPrice = totalPrice

	if err := biz.repo.HandleCancelNote(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleIngredientDetail(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleIngredientTotalAmount(ctx, mapIngredient); err != nil {
		return err
	}

	return nil
}

func handleCancelNoteId(
	gen generator.IdGenerator,
	data *cancelnotemodel.CancelNoteCreate) error {
	idCancelNote, errGenerateIdCancelNote := gen.IdProcess(data.Id)
	if errGenerateIdCancelNote != nil {
		return errGenerateIdCancelNote
	}
	data.Id = idCancelNote

	for i := range data.CancelNoteCreateDetails {
		data.CancelNoteCreateDetails[i].CancelNoteId = *idCancelNote
	}

	return nil
}

func getMapIngredientExist(data *cancelnotemodel.CancelNoteCreate) map[string]float32 {
	mapIngredientExist := make(map[string]float32)
	for _, v := range data.CancelNoteCreateDetails {
		mapIngredientExist[v.IngredientId] += v.AmountCancel
	}

	return mapIngredientExist
}
