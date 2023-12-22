package invoicerepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
)

type ListInvoiceStore interface {
	ListInvoice(
		ctx context.Context,
		filter *invoicemodel.Filter,
		propertiesContainSearchKey []string,
		paging *common.Paging,
		moreKeys ...string,
	) ([]invoicemodel.Invoice, error)
}

type listInvoiceRepo struct {
	store ListInvoiceStore
}

func NewListImportNoteRepo(store ListInvoiceStore) *listInvoiceRepo {
	return &listInvoiceRepo{store: store}
}

func (repo *listInvoiceRepo) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	result, err := repo.store.ListInvoice(
		ctx,
		filter,
		[]string{"Invoice.id"},
		paging,
		"Customer", "CreatedByUser")

	if err != nil {
		return nil, err
	}

	return result, nil
}
