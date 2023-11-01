package importnotebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/asyncjob"
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
	) (*[]importnotedetailmodel.ImportNoteDetail, error)
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
	UpdateIngredient(
		ctx context.Context,
		id string,
		data *ingredientmodel.IngredientUpdate,
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

type changeStatusImportNoteBiz struct {
	importNoteStore       ChangeStatusImportNoteStorage
	importNoteDetailStore GetImportNoteDetailStorage
	ingredientStore       UpdateAmountIngredientStorage
	ingredientDetailStore UpdateOrCreateIngredientDetailStorage
	supplierStore         UpdateDebtOfSupplierStorage
	supplierDebtStore     CreateSupplierDebtStorage
}

func NewChangeStatusImportNoteBiz(
	importNoteStore ChangeStatusImportNoteStorage,
	importNoteDetailStore GetImportNoteDetailStorage,
	ingredientStore UpdateAmountIngredientStorage,
	ingredientDetailStore UpdateOrCreateIngredientDetailStorage,
	supplierStore UpdateDebtOfSupplierStorage,
	supplierDebtStore CreateSupplierDebtStorage) *changeStatusImportNoteBiz {
	return &changeStatusImportNoteBiz{
		importNoteStore:       importNoteStore,
		importNoteDetailStore: importNoteDetailStore,
		ingredientStore:       ingredientStore,
		ingredientDetailStore: ingredientDetailStore,
		supplierStore:         supplierStore,
		supplierDebtStore:     supplierDebtStore,
	}
}

func (biz *changeStatusImportNoteBiz) ChangeStatusImportNote(
	ctx context.Context,
	importNoteId string,
	data *importnotemodel.ImportNoteUpdate) error {
	//validate data
	if err := data.Validate(); err != nil {
		return err
	}

	importNote, errCheckInProgress := biz.importNoteStore.FindImportNote(ctx, map[string]interface{}{"id": importNoteId})

	if errCheckInProgress != nil {
		return errCheckInProgress
	}

	if *importNote.Status != importnotemodel.InProgress {
		return common.ErrInternal(errors.New("import note already has been closed"))
	}

	if *data.Status == importnotemodel.Cancel {
		if err := biz.importNoteStore.UpdateImportNote(ctx, importNoteId, data); err != nil {
			return err
		}

		return nil
	}

	//get import note detail
	importNoteDetails, errGetImportNoteDetails := biz.importNoteDetailStore.FindListImportNoteDetail(
		ctx,
		map[string]interface{}{"importNoteId": importNoteId})
	if errGetImportNoteDetails != nil {
		return errGetImportNoteDetails
	}

	//handle store data
	///handle define job update status import note
	jobUpdateStatusImportNote := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.importNoteStore.UpdateImportNote(ctx, importNoteId, data)
	})

	///handle define job update debt for supplier
	supplierUpdateDebt := suppliermodel.SupplierUpdateDebt{
		Amount: &importNote.TotalPrice,
	}

	jobUpdateSupplier := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.supplierStore.UpdateSupplierDebt(ctx, importNote.SupplierId, &supplierUpdateDebt)
	})

	///handle define job create supplier debt
	debtCurrent, err := biz.supplierStore.GetDebtSupplier(
		ctx,
		importNote.SupplierId)
	if err != nil {
		return err
	}

	amountBorrow := importNote.TotalPrice
	amountLeft := *debtCurrent + amountBorrow

	idSupplierDebtCreate, errGenerateIdSupplierDebtCreate := common.GenerateId()
	if errGenerateIdSupplierDebtCreate != nil {
		return errGenerateIdSupplierDebtCreate
	}

	supplierDebtType := supplierdebtmodel.Debt
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:               idSupplierDebtCreate,
		IdSupplier:       importNote.SupplierId,
		Amount:           amountBorrow,
		AmountLeft:       amountLeft,
		SupplierDebtType: &supplierDebtType,
		CreateBy:         data.CloseBy,
	}

	jobCreateSupplierDebt := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.supplierDebtStore.CreateSupplierDebt(ctx, &supplierDebtCreate)
	})

	///handle define job update or create ingredient detail
	var ingredientDetailJobs []asyncjob.Job
	mapIngredientAmount := map[string]float32{}
	for _, v := range *importNoteDetails {
		mapIngredientAmount[v.IngredientId] += v.AmountImport
		tempJob := asyncjob.NewJob(func(
			importNoteDetail importnotedetailmodel.ImportNoteDetail,
		) func(ctx context.Context) error {
			return func(ctx context.Context) error {
				_, err := biz.ingredientDetailStore.FindIngredientDetail(
					ctx,
					map[string]interface{}{
						"ingredientId": importNoteDetail.IngredientId,
						"expiryDate":   importNoteDetail.ExpiryDate,
					},
				)
				if err != nil {
					var appErr *common.AppError
					errors.As(err, &appErr)
					if appErr != nil {
						if appErr.Key == common.ErrRecordNotFound().Key {
							dataCreate := ingredientdetailmodel.IngredientDetailCreate{
								IngredientId: importNoteDetail.IngredientId,
								ExpiryDate:   importNoteDetail.ExpiryDate,
								Amount:       importNoteDetail.AmountImport,
							}
							return biz.ingredientDetailStore.CreateIngredientDetail(
								ctx,
								&dataCreate,
							)
						}
					}
					return err
				}
				dataUpdate := ingredientdetailmodel.IngredientDetailUpdate{
					Amount: importNoteDetail.AmountImport,
				}
				return biz.ingredientDetailStore.UpdateIngredientDetail(
					ctx,
					importNoteDetail.IngredientId,
					importNoteDetail.ExpiryDate,
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
				ingredientUpdate := ingredientmodel.IngredientUpdate{Amount: amount}
				return biz.ingredientStore.UpdateIngredient(ctx, id, &ingredientUpdate)
			}
		}(value, key))
		ingredientJobs = append(ingredientJobs, tempJob)
	}

	///combine all job
	jobs := []asyncjob.Job{
		jobUpdateStatusImportNote,
		jobUpdateSupplier,
		jobCreateSupplierDebt,
	}
	jobs = append(jobs, ingredientJobs...)
	jobs = append(jobs, ingredientDetailJobs...)

	///run jobs
	group := asyncjob.NewGroup(
		false,
		jobs...)
	if err := group.Run(context.Background()); err != nil {
		return err
	}

	return nil
}
