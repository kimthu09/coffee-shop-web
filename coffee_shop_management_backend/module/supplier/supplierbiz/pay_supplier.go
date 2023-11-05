package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/component/asyncjob"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplierdebt/supplierdebtmodel"
	"context"
)

type PaySupplierStore interface {
	GetDebtSupplier(
		ctx context.Context,
		supplierId string) (*float32, error)
	UpdateSupplierDebt(
		ctx context.Context,
		id string,
		data *suppliermodel.SupplierUpdateDebt,
	) error
}

type CreateSupplierDebtStore interface {
	CreateSupplierDebt(
		ctx context.Context,
		data *supplierdebtmodel.SupplierDebtCreate) error
}

type paySupplierBiz struct {
	supplierStore     PaySupplierStore
	supplierDebtStore CreateSupplierDebtStore
	requester         common.Requester
}

func NewUpdatePayBiz(
	supplierStore PaySupplierStore,
	supplierDebtStore CreateSupplierDebtStore,
	requester common.Requester) *paySupplierBiz {
	return &paySupplierBiz{
		supplierStore:     supplierStore,
		supplierDebtStore: supplierDebtStore,
		requester:         requester}
}

func (biz *paySupplierBiz) PaySupplier(
	ctx context.Context,
	idSupplier string,
	data *suppliermodel.SupplierUpdateDebt) (*string, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	if *data.Amount <= 0 {
		return nil, suppliermodel.ErrDebtPayIsInvalid
	}

	debtCurrent, err := biz.supplierStore.GetDebtSupplier(
		ctx,
		idSupplier)
	if err != nil {
		return nil, err
	}

	amountPay := *data.Amount
	amountLeft := *debtCurrent + amountPay

	id, errGenerateId := common.GenerateId()
	if errGenerateId != nil {
		return nil, errGenerateId
	}

	debtType := enum.Pay
	supplierDebtCreate := supplierdebtmodel.SupplierDebtCreate{
		Id:         id,
		IdSupplier: idSupplier,
		Amount:     amountPay,
		AmountLeft: amountLeft,
		DebtType:   &debtType,
		CreateBy:   biz.requester.GetUserId(),
	}

	jobUpdateSupplier := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.supplierStore.UpdateSupplierDebt(ctx, idSupplier, data)
	})
	jobCreateSupplierDebt := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.supplierDebtStore.CreateSupplierDebt(ctx, &supplierDebtCreate)
	})
	group := asyncjob.NewGroup(
		false,
		jobUpdateSupplier,
		jobCreateSupplierDebt)
	if err := group.Run(context.Background()); err != nil {
		return nil, err
	}

	return &supplierDebtCreate.Id, nil
}
