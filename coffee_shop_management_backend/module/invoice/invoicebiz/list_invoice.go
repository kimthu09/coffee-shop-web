package invoicebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
)

type ListInvoiceRepo interface {
	ListInvoice(
		ctx context.Context,
		filter *invoicemodel.Filter,
		paging *common.Paging,
	) ([]invoicemodel.Invoice, error)
}

type listInvoiceBiz struct {
	repo      ListInvoiceRepo
	requester middleware.Requester
}

func NewListImportNoteBiz(
	repo ListInvoiceRepo,
	requester middleware.Requester) *listInvoiceBiz {
	return &listInvoiceBiz{repo: repo, requester: requester}
}

func (biz *listInvoiceBiz) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	if !biz.requester.IsHasFeature(common.InvoiceViewFeatureCode) {
		return nil, invoicemodel.ErrInvoiceViewNoPermission
	}

	result, err := biz.repo.ListInvoice(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
