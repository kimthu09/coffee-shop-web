package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type PaySupplierStoreRepo interface {
	GetDebtSupplier(
		ctx context.Context,
		supplierId string,
	) (*float32, error)
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate,
	) error
	UpdateDebtSupplier(
		ctx context.Context,
		supplierId string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type paySupplierBiz struct {
	gen       generator.IdGenerator
	repo      PaySupplierStoreRepo
	requester middleware.Requester
}

func NewUpdatePayBiz(
	gen generator.IdGenerator,
	repo PaySupplierStoreRepo,
	requester middleware.Requester) *paySupplierBiz {
	return &paySupplierBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *paySupplierBiz) PaySupplier(
	ctx context.Context,
	supplierId string,
	data *suppliermodel.SupplierUpdateDebt) (*string, error) {
	if !biz.requester.IsHasFeature(common.SupplierPayFeatureCode) {
		return nil, suppliermodel.ErrSupplierPayNoPermission
	}

	if err := validateSupplierUpdateDebt(data); err != nil {
		return nil, err
	}

	debtCurrent, errGetDebt := biz.repo.GetDebtSupplier(ctx, supplierId)
	if errGetDebt != nil {
		return nil, errGetDebt
	}

	supplierDebtCreate, errGetSupplierDebtCreate := getSupplierDebtCreate(
		biz.gen, supplierId, *debtCurrent, data,
	)
	if errGetSupplierDebtCreate != nil {
		return nil, errGetSupplierDebtCreate
	}

	if err := biz.repo.UpdateDebtSupplier(ctx, supplierId, data); err != nil {
		return nil, err
	}

	if err := biz.repo.CreateSupplierDebt(ctx, supplierDebtCreate); err != nil {
		return nil, err
	}

	return &supplierDebtCreate.Id, nil
}

func validateSupplierUpdateDebt(data *suppliermodel.SupplierUpdateDebt) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if *data.Amount <= 0 {
		return suppliermodel.ErrDebtPayIsInvalid
	}

	return nil
}

func getSupplierDebtCreate(
	gen generator.IdGenerator,
	supplierId string,
	currentDebt float32,
	data *suppliermodel.SupplierUpdateDebt,
) (*supplierdebtmodel.SupplierDebtCreate, error) {
	amountPay := *data.Amount
	amountLeft := currentDebt + amountPay

	id, errGenerateId := gen.GenerateId()
	if errGenerateId != nil {
		return nil, errGenerateId
	}

	debtType := enum.Pay
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         id,
		SupplierId: supplierId,
		Amount:     amountPay,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   data.CreateBy,
	}

	return &supplierDebtCreate, nil
}
