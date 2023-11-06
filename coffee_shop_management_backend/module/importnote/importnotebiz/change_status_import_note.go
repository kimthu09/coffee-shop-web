package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
)

type ChangeStatusImportNoteRepo interface {
	UpdateDebtSupplier(
		ctx context.Context,
		importNote *importnotemodel.ImportNoteUpdate) error
	CreateSupplierDebt(
		ctx context.Context,
		supplierDebtId string,
		importNote *importnotemodel.ImportNoteUpdate) error
	FindImportNote(
		ctx context.Context,
		importNoteId string,
	) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		importNoteId string,
		data *importnotemodel.ImportNoteUpdate,
	) error
	FindListImportNoteDetail(
		ctx context.Context,
		importNoteId string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
	HandleIngredientDetails(
		ctx context.Context,
		importNoteDetails []importnotedetailmodel.ImportNoteDetail,
	) error
	HandleIngredientTotalAmount(
		ctx context.Context,
		ingredientTotalAmountNeedUpdate map[string]float32,
	) error
}

type changeStatusImportNoteRepo struct {
	gen       generator.IdGenerator
	repo      ChangeStatusImportNoteRepo
	requester middleware.Requester
}

func NewChangeStatusImportNoteBiz(
	gen generator.IdGenerator,
	repo ChangeStatusImportNoteRepo,
	requester middleware.Requester) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *changeStatusImportNoteRepo) ChangeStatusImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	if !biz.requester.IsHasFeature(common.ImportNoteChangeStatusFeatureCode) {
		return importnotemodel.ErrImportNoteChangeStatusNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	importNote, errGetImportNote := biz.repo.FindImportNote(ctx, importNoteId)
	if errGetImportNote != nil {
		return errGetImportNote
	}
	data.Id = importNoteId
	data.TotalPrice = importNote.TotalPrice
	data.SupplierId = importNote.SupplierId

	if *importNote.Status != importnotemodel.InProgress {
		return importnotemodel.ErrImportNoteClosed
	}

	if *data.Status == importnotemodel.Done {
		supplierDebtId, errGenerateId := biz.gen.GenerateId()
		if errGenerateId != nil {
			return errGenerateId
		}

		if err := biz.repo.CreateSupplierDebt(ctx, supplierDebtId, data); err != nil {
			return err
		}

		if err := biz.repo.UpdateDebtSupplier(ctx, data); err != nil {
			return err
		}

		importNoteDetails, errGetImportNoteDetails := biz.repo.FindListImportNoteDetail(
			ctx,
			importNoteId)
		if errGetImportNoteDetails != nil {
			return errGetImportNoteDetails
		}

		if err := biz.repo.HandleIngredientDetails(ctx, importNoteDetails); err != nil {
			return err
		}

		mapIngredientAmount := getMapIngredientTotalAmountNeedUpdated(importNoteDetails)
		if err := biz.repo.HandleIngredientTotalAmount(ctx, mapIngredientAmount); err != nil {
			return err
		}
	}
	if err := biz.repo.UpdateImportNote(ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func getMapIngredientTotalAmountNeedUpdated(
	importNoteDetails []importnotedetailmodel.ImportNoteDetail) map[string]float32 {
	result := make(map[string]float32)
	for _, v := range importNoteDetails {
		result[v.IngredientId] += v.AmountImport
	}
	return result
}
