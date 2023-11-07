package invoicedetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"context"
)

type ListInvoiceDetailStore interface {
	ListInvoiceDetail(
		ctx context.Context,
		invoiceId string,
		paging *common.Paging,
	) ([]invoicedetailmodel.InvoiceDetail, error)
}

type listInvoiceDetailBiz struct {
	store     ListInvoiceDetailStore
	requester middleware.Requester
}

func NewListInvoiceDetailBiz(
	store ListInvoiceDetailStore,
	requester middleware.Requester) *listInvoiceDetailBiz {
	return &listInvoiceDetailBiz{store: store, requester: requester}
}

func (biz *listInvoiceDetailBiz) ListInvoiceDetail(
	ctx context.Context,
	invoiceId string,
	paging *common.Paging) ([]invoicedetailmodel.InvoiceDetail, error) {
	if !biz.requester.IsHasFeature(common.InvoiceViewFeatureCode) {
		return nil, invoicedetailmodel.ErrInvoiceDetailViewNoPermission
	}

	result, err := biz.store.ListInvoiceDetail(
		ctx,
		invoiceId,
		paging,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
