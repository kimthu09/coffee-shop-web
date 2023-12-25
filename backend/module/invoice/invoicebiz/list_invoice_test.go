package invoicebiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListInvoiceRepo struct {
	mock.Mock
}

func (m *mockListInvoiceRepo) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	args := m.Called(ctx, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]invoicemodel.Invoice), args.Error(1)
}

func TestNewListImportNoteBiz(t *testing.T) {
	type args struct {
		repo      ListInvoiceRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListInvoiceRepo)

	tests := []struct {
		name string
		args args
		want *listInvoiceBiz
	}{
		{
			name: "Create object has type ListInvoiceBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listInvoiceBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListInvoiceBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListInvoiceBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listInvoiceBiz_ListInvoice(t *testing.T) {
	type fields struct {
		repo      ListInvoiceRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *invoicemodel.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListInvoiceRepo)
	mockRequest := new(mockRequester)
	filter := invoicemodel.Filter{
		SearchKey:         "",
		DateFromCreatedAt: nil,
		DateToCreatedAt:   nil,
		MinPrice:          nil,
		MaxPrice:          nil,
		CreatedBy:         nil,
		Customer:          nil,
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	listInvoice := make([]invoicemodel.Invoice, 0)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []invoicemodel.Invoice
		wantErr bool
	}{
		{
			name: "List invoice failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listInvoice,
			wantErr: true,
		},
		{
			name: "List invoice failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListInvoice",
						context.Background(),
						&filter,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    listInvoice,
			wantErr: true,
		},
		{
			name: "List invoice successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListInvoice",
						context.Background(),
						&filter,
						&paging,
					).
					Return(listInvoice, nil).
					Once()
			},
			want:    listInvoice,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listInvoiceBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListInvoice(
				tt.args.ctx,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListInvoice() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListInvoice() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListInvoice() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
