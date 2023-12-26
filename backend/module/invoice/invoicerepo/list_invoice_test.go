package invoicerepo

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListInvoiceStore struct {
	mock.Mock
}

func (m *mockListInvoiceStore) ListInvoice(
	ctx context.Context,
	filter *invoicemodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging,
	moreKeys ...string,
) ([]invoicemodel.Invoice, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]invoicemodel.Invoice), args.Error(1)
}

func TestNewListImportNoteRepo(t *testing.T) {
	type args struct {
		store ListInvoiceStore
	}

	store := new(mockListInvoiceStore)

	tests := []struct {
		name string
		args args
		want *listInvoiceRepo
	}{
		{
			name: "Create object has type listInvoiceRepo",
			args: args{
				store: store,
			},
			want: &listInvoiceRepo{
				store: store,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListImportNoteRepo(tt.args.store)
			assert.Equal(t, tt.want, got, "NewListImportNoteRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_listInvoiceRepo_ListInvoice(t *testing.T) {
	type fields struct {
		store ListInvoiceStore
	}
	type args struct {
		ctx    context.Context
		filter *invoicemodel.Filter
		paging *common.Paging
	}

	store := new(mockListInvoiceStore)

	filterInvoice := &invoicemodel.Filter{
		SearchKey:         "invoice123",
		DateFromCreatedAt: nil,
		DateToCreatedAt:   nil,
		MinPrice:          nil,
		MaxPrice:          nil,
		CreatedBy:         nil,
		Customer:          nil,
	}

	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}

	moreKeys := []string{"Customer", "CreatedByUser"}
	listInvoices := make([]invoicemodel.Invoice, 0)

	mockErr := assert.AnError

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []invoicemodel.Invoice
		wantErr bool
	}{
		{
			name: "List invoices failed because can not get data from database",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterInvoice,
				paging: paging,
			},
			mock: func() {
				store.
					On("ListInvoice",
						context.Background(),
						filterInvoice,
						[]string{"Invoice.id"},
						paging,
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List invoices successfully",
			fields: fields{
				store: store,
			},
			args: args{
				ctx:    context.Background(),
				filter: filterInvoice,
				paging: paging,
			},
			mock: func() {
				store.
					On("ListInvoice",
						context.Background(),
						filterInvoice,
						[]string{"Invoice.id"},
						paging,
						moreKeys).
					Return(listInvoices, nil).
					Once()
			},
			want:    listInvoices,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &listInvoiceRepo{
				store: tt.fields.store,
			}

			tt.mock()

			got, err := repo.ListInvoice(tt.args.ctx, tt.args.filter, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListInvoice() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListInvoice() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
