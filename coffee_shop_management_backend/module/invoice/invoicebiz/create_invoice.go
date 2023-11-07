package invoicebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
)

type CreateInvoiceRepo interface {
	HandleCheckPermissionStatus(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	HandleData(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	CreateCustomerDebt(
		ctx context.Context,
		supplierDebtId string,
		data *invoicemodel.InvoiceCreate) error
	UpdateDebtCustomer(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate) error
	HandleInvoice(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
}

type createInvoiceBiz struct {
	gen       generator.IdGenerator
	repo      CreateInvoiceRepo
	requester middleware.Requester
}

func NewCreateInvoiceBiz(
	gen generator.IdGenerator,
	repo CreateInvoiceRepo,
	requester middleware.Requester) *createInvoiceBiz {
	return &createInvoiceBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createInvoiceBiz) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	if !biz.requester.IsHasFeature(common.InvoiceCreateFeatureCode) {
		return invoicemodel.ErrInvoiceCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInvoiceId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.HandleCheckPermissionStatus(ctx, data); err != nil {
		return err
	}

	if err := biz.repo.HandleData(ctx, data); err != nil {
		return err
	}

	if err := checkPrice(data); err != nil {
		return err
	}

	if data.CustomerId != nil {
		customerDebtId, errGenerateId := biz.gen.GenerateId()
		if errGenerateId != nil {
			return errGenerateId
		}

		if err := biz.repo.CreateCustomerDebt(ctx, customerDebtId, data); err != nil {
			return err
		}

		if err := biz.repo.UpdateDebtCustomer(ctx, data); err != nil {
			return err
		}
	}

	if err := biz.repo.HandleInvoice(ctx, data); err != nil {
		return err
	}
	return nil
}

func handleInvoiceId(gen generator.IdGenerator, data *invoicemodel.InvoiceCreate) error {
	idInvoice, err := gen.GenerateId()
	if err != nil {
		return err
	}
	data.Id = idInvoice

	for i := range data.InvoiceDetails {
		data.InvoiceDetails[i].InvoiceId = idInvoice
	}

	return err
}

func checkPrice(data *invoicemodel.InvoiceCreate) error {
	data.AmountDebt = data.TotalPrice - data.AmountReceived
	if data.CustomerId == nil && data.AmountDebt > 0 {
		return invoicemodel.ErrInvoiceNotHaveCustomerForDebt
	}
	if common.ValidateNegativeNumber(data.AmountDebt) {
		return invoicemodel.ErrInvoiceAmountDebtIsNegativeNumber
	}
	return nil
}
