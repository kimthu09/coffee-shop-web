package importnoterepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
	"errors"
)

type ChangeStatusImportNoteStorage interface {
	FindImportNote(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*importnotemodel.ImportNote, error)
	UpdateImportNote(
		ctx context.Context,
		id string,
		data *importnotemodel.ImportNoteUpdate,
	) error
}

type GetImportNoteDetailStorage interface {
	FindListImportNoteDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type UpdateOrCreateIngredientDetailStorage interface {
	FindIngredientDetail(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ingredientdetailmodel.IngredientDetail, error)
	UpdateIngredientDetail(
		ctx context.Context,
		ingredientId string,
		expiryDate string,
		data *ingredientdetailmodel.IngredientDetailUpdate,
	) error
	CreateIngredientDetail(
		ctx context.Context,
		data *ingredientdetailmodel.IngredientDetailCreate,
	) error
}

type UpdateAmountIngredientStorage interface {
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
}

type UpdateDebtOfSupplierStorage interface {
	GetDebtSupplier(
		ctx context.Context,
		supplierId string,
	) (*float32, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type CreateSupplierDebtStorage interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
}

type changeStatusImportNoteRepo struct {
	importNoteStore       ChangeStatusImportNoteStorage
	importNoteDetailStore GetImportNoteDetailStorage
	ingredientStore       UpdateAmountIngredientStorage
	ingredientDetailStore UpdateOrCreateIngredientDetailStorage
	supplierStore         UpdateDebtOfSupplierStorage
	supplierDebtStore     CreateSupplierDebtStorage
}

func NewChangeStatusImportNoteRepo(
	importNoteStore ChangeStatusImportNoteStorage,
	importNoteDetailStore GetImportNoteDetailStorage,
	ingredientStore UpdateAmountIngredientStorage,
	ingredientDetailStore UpdateOrCreateIngredientDetailStorage,
	supplierStore UpdateDebtOfSupplierStorage,
	supplierDebtStore CreateSupplierDebtStorage) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		ingredientStore:       ingredientStore,
		ingredientDetailStore: ingredientDetailStore,
		supplierStore:         supplierStore,
		supplierDebtStore:     supplierDebtStore,
	}
}

func (repo *changeStatusImportNoteRepo) FindImportNote(
	ctx context.Context,
	importNoteId string) (*importnotemodel.ImportNote, error) {
	importNote, err := repo.importNoteStore.FindImportNote(ctx, map[string]interface{}{"id": importNoteId})
	if err != nil {
		return nil, err
	}
	return importNote, nil
}

func (repo *changeStatusImportNoteRepo) UpdateImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	if err := repo.importNoteStore.UpdateImportNote(ctx, importNoteId, data); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) CreateSupplierDebt(
	ctx context.Context,
	supplierDebtId string,
	importNote *importnotemodel.ImportNoteUpdate) error {
	debtCurrent, err := repo.supplierStore.GetDebtSupplier(
		ctx,
		importNote.SupplierId)
	if err != nil {
		return err
	}

	amountBorrow := importNote.TotalPrice
	amountLeft := *debtCurrent + amountBorrow

	debtType := enum.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         supplierDebtId,
		SupplierId: importNote.SupplierId,
		Amount:     amountBorrow,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   importNote.CloseBy,
	}

	if err := repo.supplierDebtStore.CreateSupplierDebt(
		ctx, &supplierDebtCreate,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) UpdateDebtSupplier(
	ctx context.Context,
	importNote *importnotemodel.ImportNoteUpdate) error {
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount: &importNote.TotalPrice,
	}
	if err := repo.supplierStore.UpdateSupplierDebt(
		ctx, importNote.SupplierId, &supplierUpdateDebt,
	); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) FindListImportNoteDetail(
	ctx context.Context,
	importNoteId string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	importNoteDetails, errGetImportNoteDetails := repo.importNoteDetailStore.FindListImportNoteDetail(
		ctx,
		map[string]interface{}{"importNoteId": importNoteId})
	if errGetImportNoteDetails != nil {
		return nil, errGetImportNoteDetails
	}
	return importNoteDetails, nil
}

func (repo *changeStatusImportNoteRepo) HandleIngredientDetails(
	ctx context.Context,
	importNoteDetails []importnotedetailmodel.ImportNoteDetail) error {
	var updatedIngredientDetails []ingredientdetailmodel.IngredientDetailUpdate
	var createdIngredientDetails []ingredientdetailmodel.IngredientDetailCreate

	for _, v := range importNoteDetails {
		_, err := repo.ingredientDetailStore.FindIngredientDetail(
			ctx,
			map[string]interface{}{
				"ingredientId": v.IngredientId,
				"expiryDate":   v.ExpiryDate,
			},
		)

		if err != nil {
			var appErr *common.AppError
			errors.As(err, &appErr)
			if appErr != nil {
				if appErr.Key == common.ErrRecordNotFound().Key {
					dataCreate := ingredientdetailmodel.IngredientDetailCreate{
						IngredientId: v.IngredientId,
						ExpiryDate:   v.ExpiryDate,
						Amount:       v.AmountImport,
					}
					createdIngredientDetails = append(
						createdIngredientDetails, dataCreate,
					)
				}
			} else {
				return err
			}
		} else {
			dataUpdate := ingredientdetailmodel.IngredientDetailUpdate{
				IngredientId: v.IngredientId,
				ExpiryDate:   v.ExpiryDate,
				Amount:       v.AmountImport,
			}
			updatedIngredientDetails = append(updatedIngredientDetails, dataUpdate)
		}

	}

	if err := repo.updateIngredientDetails(ctx, updatedIngredientDetails); err != nil {
		return err
	}

	if err := repo.createIngredientDetails(ctx, createdIngredientDetails); err != nil {
		return err
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) updateIngredientDetails(
	ctx context.Context,
	updatedIngredientDetails []ingredientdetailmodel.IngredientDetailUpdate) error {
	for _, v := range updatedIngredientDetails {
		if err := repo.ingredientDetailStore.UpdateIngredientDetail(
			ctx, v.IngredientId, v.ExpiryDate, &v,
		); err != nil {
			return err
		}
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) createIngredientDetails(
	ctx context.Context,
	createdIngredientDetails []ingredientdetailmodel.IngredientDetailCreate) error {
	for _, v := range createdIngredientDetails {
		if err := repo.ingredientDetailStore.CreateIngredientDetail(
			ctx, &v,
		); err != nil {
			return err
		}
	}
	return nil
}

func (repo *changeStatusImportNoteRepo) HandleIngredientTotalAmount(
	ctx context.Context,
	ingredientTotalAmountNeedUpdate map[string]float32) error {
	for key, value := range ingredientTotalAmountNeedUpdate {
		ingredientUpdate := ingredientmodel.IngredientUpdateAmount{Amount: value}
		if err := repo.ingredientStore.UpdateAmountIngredient(
			ctx, key, &ingredientUpdate,
		); err != nil {
			return err
		}
	}
	return nil
}
