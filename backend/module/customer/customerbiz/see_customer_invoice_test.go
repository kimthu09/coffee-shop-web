package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeCustomerInvoiceRepo struct {
	mock.Mock
}

func (m *mockSeeCustomerInvoiceRepo) SeeCustomerInvoice(
	ctx context.Context,
	customerId string,
	filter *customermodel.FilterInvoice,
	paging *common.Paging) ([]invoicemodel.Invoice, error) {
	args := m.Called(ctx, customerId, filter, paging)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]invoicemodel.Invoice), args.Error(1)
}

func TestNewSeeCustomerInvoiceBiz(t *testing.T) {
	type args struct {
		repo      SeeCustomerInvoiceRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeCustomerInvoiceRepo)

	tests := []struct {
		name string
		args args
		want *seeCustomerInvoiceBiz
	}{
		{
			name: "Create object has type SeeCustomerInvoiceBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeCustomerInvoiceBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeCustomerInvoiceBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeCustomerInvoiceBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeCustomerInvoiceBiz_SeeCustomerInvoice(t *testing.T) {
	type fields struct {
		repo      SeeCustomerInvoiceRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		customerId string
		filter     *customermodel.FilterInvoice
		paging     *common.Paging
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeCustomerInvoiceRepo)
	customerId := mock.Anything
	date := int64(123)
	filterInvoice := customermodel.FilterInvoice{
		DateFrom: &date,
		DateTo:   &date,
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
		Total: 12,
	}
	invoices := []invoicemodel.Invoice{
		{
			Id:             mock.Anything,
			CustomerId:     customerId,
			Customer:       invoicemodel.SimpleCustomer{Id: customerId, Name: mock.Anything},
			TotalPrice:     0,
			AmountReceived: 0,
			CreatedBy:      mock.Anything,
			CreatedAt:      nil,
		},
	}
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
			name: "See customer invoice failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				filter:     &filterInvoice,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See customer invoice failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				filter:     &filterInvoice,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeCustomerInvoice",
						context.Background(),
						customerId,
						&filterInvoice,
						&paging,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See customer invoice successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				filter:     &filterInvoice,
				paging:     &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeCustomerInvoice",
						context.Background(),
						customerId,
						&filterInvoice,
						&paging,
					).
					Return(invoices, nil).
					Once()
			},
			want:    invoices,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeCustomerInvoiceBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeCustomerInvoice(
				tt.args.ctx,
				tt.args.customerId,
				tt.args.filter,
				tt.args.paging,
			)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeCustomerInvoice() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeCustomerInvoice() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeCustomerInvoice() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
