package importnoterepo

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type ChangeStatusImportNoteStore interface {
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

type GetImportNoteDetailStore interface {
	FindListImportNoteDetail(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) ([]importnotedetailmodel.ImportNoteDetail, error)
}

type UpdateAmountIngredientStore interface {
	UpdateAmountIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdateAmount,
	) error
}

type UpdateDebtOfSupplierStore interface {
	FindSupplier(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string) (*suppliermodel.Supplier, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
}

type changeStatusImportNoteRepo struct {
	importNoteStore       ChangeStatusImportNoteStore
	importNoteDetailStore GetImportNoteDetailStore
	ingredientStore       UpdateAmountIngredientStore
	supplierStore         UpdateDebtOfSupplierStore
	supplierDebtStore     CreateSupplierDebtStore
}

func NewChangeStatusImportNoteRepo(
	importNoteStore ChangeStatusImportNoteStore,
	importNoteDetailStore GetImportNoteDetailStore,
	ingredientStore UpdateAmountIngredientStore,
	supplierStore UpdateDebtOfSupplierStore,
	supplierDebtStore CreateSupplierDebtStore) *changeStatusImportNoteRepo {
	return &changeStatusImportNoteRepo{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		ingredientStore:       ingredientStore,
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
	supplier, err := repo.supplierStore.FindSupplier(
		ctx,
		map[string]interface{}{"id": importNote.SupplierId})
	if err != nil {
		return err
	}

	amountBorrow := importNote.TotalPrice
	amountLeft := supplier.Debt - amountBorrow

	debtType := enum.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         supplierDebtId,
		SupplierId: importNote.SupplierId,
		Amount:     -amountBorrow,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreatedBy:  importNote.ClosedBy,
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
	amount := -importNote.TotalPrice
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount: &amount,
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

func (repo *changeStatusImportNoteRepo) HandleIngredient(
	ctx context.Context,
	ingredientTotalAmountNeedUpdate map[string]int) error {
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
