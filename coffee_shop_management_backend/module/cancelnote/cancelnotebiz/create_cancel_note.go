package cancelnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/asyncjob"
	"coffee_shop_management_backend/module/cancelnote/cancelnotemodel"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
)

type CreateCancelNoteStorage interface {
	CreateCancelNote(
		ctx context.Context,
		data *cancelnotemodel.CancelNoteCreate,
	) error
}

type CreateCancelNoteDetailStorage interface {
	CreateListCancelNoteDetail(
		ctx context.Context,
		data []cancelnotedetailmodel.CancelNoteDetailCreate,
	) error
}

type UpdateIngredientStorage interface {
	GetPriceIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*float32, error)
	UpdateIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdate,
	) error
}

type UpdateIngredientDetailStorage interface {
	UpdateIngredientDetail(
		ctx context.Context,
		ingredientId string,
		expiryDate string,
		data *ingredientdetailmodel.IngredientDetailUpdate,
	) error
}

type createCancelNoteBiz struct {
	cancelNoteStore       CreateCancelNoteStorage
	cancelNoteDetailStore CreateCancelNoteDetailStorage
	ingredientStore       UpdateIngredientStorage
	ingredientDetailStore UpdateIngredientDetailStorage
}

func NewCreateCancelNoteBiz(
	cancelNoteStore CreateCancelNoteStorage,
	cancelNoteDetailStore CreateCancelNoteDetailStorage,
	ingredientStore UpdateIngredientStorage,
	ingredientDetailStore UpdateIngredientDetailStorage) *createCancelNoteBiz {
	return &createCancelNoteBiz{
		cancelNoteStore:       cancelNoteStore,
		cancelNoteDetailStore: cancelNoteDetailStore,
		ingredientStore:       ingredientStore,
		ingredientDetailStore: ingredientDetailStore,
	}
}

func (biz *createCancelNoteBiz) CreateCancelNote(
	ctx context.Context,
	data *cancelnotemodel.CancelNoteCreate) error {

	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	//handle id import note
	idCancelNote, errGenerateIdCancelNote := common.GenerateId()
	if errGenerateIdCancelNote != nil {
		return errGenerateIdCancelNote
	}

	data.Id = idCancelNote

	//handle export detail
	mapIngredient := map[string]float32{}

	for i, ingredientDetail := range data.CancelNoteCreateDetails {
		data.CancelNoteCreateDetails[i].CancelNoteId = idCancelNote

		mapIngredient[ingredientDetail.IngredientId] +=
			ingredientDetail.AmountCancel
	}

	//handle totalPrice
	var totalPrice float32 = 0
	for ingredientId, totalAmountOfIngredientId := range mapIngredient {
		priceIngredient, err := biz.ingredientStore.GetPriceIngredient(
			ctx,
			map[string]interface{}{"id": ingredientId},
		)
		if err != nil {
			return err
		}
		totalPrice += *priceIngredient * totalAmountOfIngredientId
	}
	data.TotalPrice = totalPrice

	//handle store data
	///handle define job create export note
	jobCreateCancelNote := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.cancelNoteStore.CreateCancelNote(ctx, data)
	})

	///handle define job create import note detail
	jobCreateCancelNoteDetail := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.cancelNoteDetailStore.CreateListCancelNoteDetail(ctx, data.CancelNoteCreateDetails)
	})

	///handle define job update ingredient detail
	var ingredientDetailJobs []asyncjob.Job
	mapIngredientAmount := map[string]float32{}
	for _, v := range data.CancelNoteCreateDetails {
		mapIngredientAmount[v.IngredientId] += v.AmountCancel
		tempJob := asyncjob.NewJob(func(
			cancelNoteDetailCreate cancelnotedetailmodel.CancelNoteDetailCreate,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				dataUpdate := ingredientdetailmodel.IngredientDetailUpdate{
					Amount: -cancelNoteDetailCreate.AmountCancel,
				}
				return biz.ingredientDetailStore.UpdateIngredientDetail(
					ctx,
					cancelNoteDetailCreate.IngredientId,
					cancelNoteDetailCreate.ExpiryDate,
					&dataUpdate,
				)
			}
		}(v))
		ingredientDetailJobs = append(ingredientDetailJobs, tempJob)
	}

	///handle define job update total amount for ingredient
	var ingredientJobs []asyncjob.Job
	for key, value := range mapIngredientAmount {
		tempJob := asyncjob.NewJob(func(amount float32, id string) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				ingredientUpdate := ingredientmodel.IngredientUpdate{Amount: -amount}
				return biz.ingredientStore.UpdateIngredient(ctx, id, &ingredientUpdate)
			}
		}(value, key))
		ingredientJobs = append(ingredientJobs, tempJob)
	}

	///combine all job
	jobs := []asyncjob.Job{
		jobCreateCancelNote,
		jobCreateCancelNoteDetail,
	}
	jobs = append(jobs, ingredientJobs...)
	jobs = append(jobs, ingredientDetailJobs...)

	///run jobs
	group := asyncjob.NewGroup(
		false,
		jobs...,
	)
	if err := group.Run(context.Background()); err != nil {
		return err
	}

	return nil
}
