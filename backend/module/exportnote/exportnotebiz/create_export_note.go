package exportnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
)

type CreateExportNoteRepo interface {
	HandleExportNote(
		ctx context.Context,
		data *exportnotemodel.ExportNoteCreate,
	) error
	HandleIngredientTotalAmount(
		ctx context.Context,
		exportNoteId string,
		ingredientTotalAmountNeedUpdate map[string]int,
	) error
}

type createExportNoteBiz struct {
	gen       generator.IdGenerator
	repo      CreateExportNoteRepo
	requester middleware.Requester
}

func NewCreateExportNoteBiz(
	gen generator.IdGenerator,
	repo CreateExportNoteRepo,
	requester middleware.Requester) *createExportNoteBiz {
	return &createExportNoteBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createExportNoteBiz) CreateExportNote(
	ctx context.Context,
	data *exportnotemodel.ExportNoteCreate) error {
	if !biz.requester.IsHasFeature(common.ExportNoteCreateFeatureCode) {
		return exportnotemodel.ErrExportNoteCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleExportNoteId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleExportNote(ctx, data); err != nil {
		return err
	}

	mapIngredient := getMapIngredientExist(data)
	if err := biz.repo.HandleIngredientTotalAmount(
		ctx, *data.Id, mapIngredient); err != nil {
		return err
	}

	return nil
}

func handleExportNoteId(
	gen generator.IdGenerator,
	data *exportnotemodel.ExportNoteCreate) error {
	idCancelNote, errGenerateIdCancelNote := gen.IdProcess(data.Id)
	if errGenerateIdCancelNote != nil {
		return errGenerateIdCancelNote
	}
	data.Id = idCancelNote

	for i := range data.ExportNoteDetails {
		data.ExportNoteDetails[i].ExportNoteId = *idCancelNote
	}

	return nil
}

func getMapIngredientExist(data *exportnotemodel.ExportNoteCreate) map[string]int {
	mapIngredientExist := make(map[string]int)
	for _, v := range data.ExportNoteDetails {
		mapIngredientExist[v.IngredientId] += v.AmountExport
	}

	return mapIngredientExist
}
