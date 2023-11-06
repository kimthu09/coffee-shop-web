package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
	"errors"
)

type ChangeStatusImportNoteRepo interface {
	HandleSupplier(
		ctx context.Context,
		data *importnotemodel.ImportNoteUpdate,
	) error
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
	repo ChangeStatusImportNoteRepo
}

func NewChangeStatusImportNoteBiz(
	repo ChangeStatusImportNoteRepo) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		repo: repo,
	}
}

func (biz *changeStatusImportNoteRepo) ChangeStatusImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	importNote, errCheckInProgress := biz.repo.FindImportNote(ctx, importNoteId)
	if errCheckInProgress != nil {
		return errCheckInProgress
	}
	data.Id = importNoteId
	data.TotalPrice = importNote.TotalPrice
	data.SupplierId = importNote.SupplierId

	if *importNote.Status != importnotemodel.InProgress {
		return common.ErrInternal(errors.New("import note already has been closed"))
	}

	if *data.Status == importnotemodel.Done {
		importNoteDetails, errGetImportNoteDetails := biz.repo.FindListImportNoteDetail(
			ctx,
			importNoteId)
		if errGetImportNoteDetails != nil {
			return errGetImportNoteDetails
		}

		if err := biz.repo.HandleSupplier(ctx, data); err != nil {
			return err
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
