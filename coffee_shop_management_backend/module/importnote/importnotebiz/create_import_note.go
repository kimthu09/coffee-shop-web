package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/asyncjob"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

type CreateImportNoteStorage interface {
	CreateImportNote(
		ctx context.Context,
		data *importnotemodel.ImportNoteCreate,
	) error
}

type CreateImportNoteDetailStorage interface {
	CreateListImportNoteDetail(
		ctx context.Context,
		data []importnotedetailmodel.ImportNoteDetailCreate,
	) error
}

type GetTotalPriceIngredientStorage interface {
	GetPriceIngredient(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*float32, error)
}

type createImportNoteBiz struct {
	importNoteStore              CreateImportNoteStorage
	importNoteDetailStore        CreateImportNoteDetailStorage
	getTotalPriceIngredientStore GetTotalPriceIngredientStorage
}

func NewCreateImportNoteBiz(
	importNoteStore CreateImportNoteStorage,
	importNoteDetailStore CreateImportNoteDetailStorage,
	getTotalPriceIngredientStore GetTotalPriceIngredientStorage) *createImportNoteBiz {
	return &createImportNoteBiz{
		importNoteStore:              importNoteStore,
		importNoteDetailStore:        importNoteDetailStore,
		getTotalPriceIngredientStore: getTotalPriceIngredientStore,
	}
}

func (biz *createImportNoteBiz) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ImportNoteCreate) error {

	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	//handle id import note
	idImportNote, errGenerateIdImportNote := common.GenerateId()
	if errGenerateIdImportNote != nil {
		return errGenerateIdImportNote
	}

	data.Id = idImportNote

	//handle import detail
	mapIngredient := map[string]float32{}

	for i, ingredientDetail := range data.ImportNoteDetails {
		data.ImportNoteDetails[i].ImportNoteId = idImportNote

		mapIngredient[ingredientDetail.IngredientId] +=
			ingredientDetail.AmountImport
	}

	//handle totalPrice
	var totalPrice float32 = 0
	for ingredientId, totalAmountOfIngredientId := range mapIngredient {
		priceIngredient, err := biz.getTotalPriceIngredientStore.GetPriceIngredient(
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
	///handle define job create import note
	jobCreateImportNote := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.importNoteStore.CreateImportNote(ctx, data)
	})

	///handle define job create import note detail
	jobCreateImportNoteDetail := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.importNoteDetailStore.CreateListImportNoteDetail(ctx, data.ImportNoteDetails)
	})

	///run jobs
	group := asyncjob.NewGroup(
		false,
		jobCreateImportNote,
		jobCreateImportNoteDetail)
	if err := group.Run(context.Background()); err != nil {
		return err
	}

	return nil
}
