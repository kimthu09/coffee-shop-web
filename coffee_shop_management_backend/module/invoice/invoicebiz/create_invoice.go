package invoicebiz

import (
	"coffee_shop_management_backend/common"
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
	HandleCustomer(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
	HandleInvoice(
		ctx context.Context,
		data *invoicemodel.InvoiceCreate,
	) error
}

type createInvoiceBiz struct {
	repo CreateInvoiceRepo
}

func NewCreateInvoiceBiz(
	repo CreateInvoiceRepo) *createInvoiceBiz {
	return &createInvoiceBiz{
		repo: repo,
	}
}

func (biz *createInvoiceBiz) CreateInvoice(
	ctx context.Context,
	data *invoicemodel.InvoiceCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleInvoiceId(data); err != nil {
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
		if err := biz.repo.HandleCustomer(ctx, data); err != nil {
			return err
		}
	}

	if err := biz.repo.HandleInvoice(ctx, data); err != nil {
		return err
	}
	return nil
}

func handleInvoiceId(data *invoicemodel.InvoiceCreate) error {
	idInvoice, err := common.GenerateId()
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
