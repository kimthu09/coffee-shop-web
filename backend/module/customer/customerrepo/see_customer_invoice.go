package customerrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
)

type ListCustomerInvoiceStore interface {
	ListAllInvoiceByCustomer(
		ctx context.Context,
		customerId string,
		filter *customermodel.FilterInvoice,
		paging *common.Paging,
		moreKeys ...string) ([]invoicemodel.Invoice, error)
}

type seeCustomerInvoiceRepo struct {
	invoiceStore ListCustomerInvoiceStore
}

func NewSeeCustomerInvoiceRepo(
	invoiceStore ListCustomerInvoiceStore) *seeCustomerInvoiceRepo {
	return &seeCustomerInvoiceRepo{
		invoiceStore: invoiceStore,
	}
}

func (biz *seeCustomerInvoiceRepo) SeeCustomerInvoice(
	ctx context.Context,
	customerId string,
	filter *customermodel.FilterInvoice,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	invoices, errInvoices := biz.invoiceStore.ListAllInvoiceByCustomer(
		ctx,
		customerId,
		filter,
		paging,
		"CreatedByUser",
	)
	if errInvoices != nil {
		return nil, errInvoices
	}

	return invoices, nil
}
