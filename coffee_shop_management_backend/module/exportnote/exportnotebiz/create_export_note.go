package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/asyncjob"
	importnotemodel "coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"context"
)

type CreateExportNoteStorage interface {
	CreateExportNote(
		ctx context.Context,
		data *importnotemodel.ExportNoteCreate,
	) error
}

type CreateExportNoteDetailStorage interface {
	CreateListExportNoteDetail(
		ctx context.Context,
		data []exportnotedetailmodel.ExportNoteDetailCreate,
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

type createExportNoteBiz struct {
	exportNoteStore       CreateExportNoteStorage
	exportNoteDetailStore CreateExportNoteDetailStorage
	ingredientStore       UpdateIngredientStorage
	ingredientDetailStore UpdateIngredientDetailStorage
}

func NewCreateExportNoteBiz(
	exportNoteStore CreateExportNoteStorage,
	exportNoteDetailStore CreateExportNoteDetailStorage,
	ingredientStore UpdateIngredientStorage,
	ingredientDetailStore UpdateIngredientDetailStorage) *createExportNoteBiz {
	return &createExportNoteBiz{
		exportNoteStore:       exportNoteStore,
		exportNoteDetailStore: exportNoteDetailStore,
		ingredientStore:       ingredientStore,
		ingredientDetailStore: ingredientDetailStore,
	}
}

func (biz *createExportNoteBiz) CreateExportNote(
	ctx context.Context,
	data *importnotemodel.ExportNoteCreate) error {

	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	//handle id import note
	idExportNote, errGenerateIdExportNote := common.GenerateId()
	if errGenerateIdExportNote != nil {
		return errGenerateIdExportNote
	}

	data.Id = idExportNote

	//handle export detail
	mapIngredient := map[string]float32{}

	for i, ingredientDetail := range data.ExportNoteDetails {
		data.ExportNoteDetails[i].ExportNoteId = idExportNote

		mapIngredient[ingredientDetail.IngredientId] +=
			ingredientDetail.AmountExport
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
	jobCreateExportNote := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.exportNoteStore.CreateExportNote(ctx, data)
	})

	///handle define job create import note detail
	jobCreateExportNoteDetail := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.exportNoteDetailStore.CreateListExportNoteDetail(ctx, data.ExportNoteDetails)
	})

	///handle define job update ingredient detail
	var ingredientDetailJobs []asyncjob.Job
	mapIngredientAmount := map[string]float32{}
	for _, v := range data.ExportNoteDetails {
		mapIngredientAmount[v.IngredientId] += v.AmountExport
		tempJob := asyncjob.NewJob(func(
			exportNoteDetailCreate exportnotedetailmodel.ExportNoteDetailCreate,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				dataUpdate := ingredientdetailmodel.IngredientDetailUpdate{
					Amount: -exportNoteDetailCreate.AmountExport,
				}
				return biz.ingredientDetailStore.UpdateIngredientDetail(
					ctx,
					exportNoteDetailCreate.IngredientId,
					exportNoteDetailCreate.ExpiryDate,
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
		jobCreateExportNote,
		jobCreateExportNoteDetail,
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
