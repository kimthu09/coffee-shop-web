package customerrepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/customer/customermodel"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockListCustomerInvoiceStore struct {
	mock.Mock
}

func (m *MockListCustomerInvoiceStore) ListAllInvoiceByCustomer(
	ctx context.Context,
	customerId string,
	filter *customermodel.FilterInvoice,
	paging *common.Paging,
	moreKeys ...string) ([]invoicemodel.Invoice, error) {
	args := m.Called(ctx, customerId, filter, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]invoicemodel.Invoice), args.Error(1)
}

func TestNewSeeCustomerInvoiceRepo(t *testing.T) {
	type args struct {
		invoiceStore ListCustomerInvoiceStore
	}

	store := new(MockListCustomerInvoiceStore)

	tests := []struct {
		name string
		args args
		want *seeCustomerInvoiceRepo
	}{
		{
			name: "Create object has type NewSeeCustomerInvoiceRepo",
			args: args{
				invoiceStore: store,
			},
			want: &seeCustomerInvoiceRepo{
				invoiceStore: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeCustomerInvoiceRepo(tt.args.invoiceStore)
			assert.Equal(t, tt.want, got, "NewSeeCustomerInvoiceRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeCustomerInvoiceRepo_SeeCustomerInvoice(t *testing.T) {
	type fields struct {
		invoiceStore ListCustomerInvoiceStore
	}
	type args struct {
		ctx        context.Context
		customerId string
		filter     *customermodel.FilterInvoice
		paging     *common.Paging
	}

	store := new(MockListCustomerInvoiceStore)

	customerId := "123"
	filter := &customermodel.FilterInvoice{}
	paging := &common.Paging{
		Page:  1,
		Limit: 10,
		Total: 12,
	}
	moreKeys := []string{"CreatedByUser"}

	mockErr := errors.New(mock.Anything)
	mockInvoices := []invoicemodel.Invoice{
		{
			Id:             "invoice001",
			CustomerId:     customerId,
			TotalPrice:     100,
			AmountReceived: 50,
			CreatedBy:      "user001",
			CreatedAt:      &time.Time{},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []invoicemodel.Invoice
		wantErr bool
	}{
		{
			name: "See customer invoice failed because can not get data from database",
			fields: fields{
				invoiceStore: store,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				filter:     filter,
				paging:     paging,
			},
			mock: func() {
				store.
					On("ListAllInvoiceByCustomer",
						context.Background(),
						customerId,
						filter,
						paging,
						moreKeys,
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
				invoiceStore: store,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
				filter:     filter,
				paging:     paging,
			},
			mock: func() {
				store.
					On("ListAllInvoiceByCustomer",
						context.Background(),
						customerId,
						filter,
						paging,
						moreKeys,
					).
					Return(mockInvoices, nil).
					Once()
			},
			want:    mockInvoices,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeCustomerInvoiceRepo{
				invoiceStore: tt.fields.invoiceStore,
			}

			tt.mock()

			got, err := biz.SeeCustomerInvoice(
				tt.args.ctx,
				tt.args.customerId,
				tt.args.filter,
				tt.args.paging,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeCustomerInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeCustomerInvoice() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "SeeCustomerInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
