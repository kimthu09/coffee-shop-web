package invoicerepo

import (
	"coffee_shop_management_backend/module/invoice/invoicemodel"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockSeeInvoiceDetailStore struct {
	mock.Mock
}

func (m *mockSeeInvoiceDetailStore) ListInvoiceDetail(
	ctx context.Context,
	invoiceId string) ([]invoicedetailmodel.InvoiceDetail, error) {
	args := m.Called(ctx, invoiceId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]invoicedetailmodel.InvoiceDetail), args.Error(1)
}

type mockFindInvoiceStore struct {
	mock.Mock
}

func (m *mockFindInvoiceStore) FindInvoice(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*invoicemodel.Invoice, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*invoicemodel.Invoice), args.Error(1)
}

func TestNewSeeInvoiceDetailRepo(t *testing.T) {
	type args struct {
		invoiceDetailStore SeeInvoiceDetailStore
		invoiceStore       FindInvoiceStore
	}

	mockInvoiceDetail := new(mockSeeInvoiceDetailStore)
	mockInvoice := new(mockFindInvoiceStore)

	tests := []struct {
		name string
		args args
		want *seeInvoiceDetailRepo
	}{
		{
			name: "Create object has type seeInvoiceDetailRepo",
			args: args{
				invoiceDetailStore: mockInvoiceDetail,
				invoiceStore:       mockInvoice,
			},
			want: &seeInvoiceDetailRepo{
				invoiceDetailStore: mockInvoiceDetail,
				invoiceStore:       mockInvoice,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeInvoiceDetailRepo(tt.args.invoiceDetailStore, tt.args.invoiceStore)
			assert.Equal(t, tt.want, got, "NewSeeInvoiceDetailRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_seeInvoiceDetailRepo_SeeInvoiceDetail(t *testing.T) {
	type fields struct {
		invoiceDetailStore SeeInvoiceDetailStore
		invoiceStore       FindInvoiceStore
	}
	type args struct {
		ctx       context.Context
		invoiceId string
	}

	mockInvoiceDetail := new(mockSeeInvoiceDetailStore)
	mockInvoice := new(mockFindInvoiceStore)
	invoiceId := "invoice123"
	invoice := invoicemodel.Invoice{
		Id:         invoiceId,
		CustomerId: "customer123",
		Customer: invoicemodel.SimpleCustomer{
			Id:    "customer123",
			Name:  "Nguyễn Văn A",
			Phone: "0123456789",
		},
		TotalPrice:          100.0,
		AmountReceived:      100,
		AmountPriceUsePoint: 0,
		CreatedBy:           "user123",
		CreatedByUser: usermodel.SimpleUser{
			Id:   "user123",
			Name: "Nguyễn Văn B",
		},
		CreatedAt: &time.Time{},
	}
	details := []invoicedetailmodel.InvoiceDetail{
		{
			InvoiceId:   invoiceId,
			FoodId:      "food001",
			SizeName:    "S",
			Amount:      10,
			UnitPrice:   100,
			Description: "",
		},
	}
	finalInvoice := invoicemodel.Invoice{
		Id:         invoiceId,
		CustomerId: "customer123",
		Customer: invoicemodel.SimpleCustomer{
			Id:    "customer123",
			Name:  "Nguyễn Văn A",
			Phone: "0123456789",
		},
		TotalPrice:          100.0,
		AmountReceived:      100,
		AmountPriceUsePoint: 0,
		CreatedBy:           "user123",
		CreatedByUser: usermodel.SimpleUser{
			Id:   "user123",
			Name: "Nguyễn Văn B",
		},
		CreatedAt: &time.Time{},
		Details: []invoicedetailmodel.InvoiceDetail{
			{
				InvoiceId:   invoiceId,
				FoodId:      "food001",
				SizeName:    "S",
				Amount:      10,
				UnitPrice:   100,
				Description: "",
			},
		},
	}
	moreKeys := []string{"Customer", "CreatedByUser"}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *invoicemodel.Invoice
		wantErr bool
	}{
		{
			name: "See invoice details failed because can not find invoice",
			fields: fields{
				invoiceDetailStore: mockInvoiceDetail,
				invoiceStore:       mockInvoice,
			},
			args: args{
				ctx:       context.Background(),
				invoiceId: invoiceId,
			},
			mock: func() {
				mockInvoice.
					On("FindInvoice",
						context.Background(),
						map[string]interface{}{"id": invoiceId},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    &finalInvoice,
			wantErr: true,
		},
		{
			name: "See invoice details failed because can not get invoice details",
			fields: fields{
				invoiceDetailStore: mockInvoiceDetail,
				invoiceStore:       mockInvoice,
			},
			args: args{
				ctx:       context.Background(),
				invoiceId: invoiceId,
			},
			mock: func() {
				mockInvoice.
					On("FindInvoice",
						context.Background(),
						map[string]interface{}{"id": invoiceId},
						moreKeys).
					Return(&invoice, nil).
					Once()

				mockInvoiceDetail.
					On("ListInvoiceDetail",
						context.Background(),
						invoiceId).
					Return(nil, mockErr).
					Once()
			},
			want:    &finalInvoice,
			wantErr: true,
		},
		{
			name: "See invoice details successfully",
			fields: fields{
				invoiceDetailStore: mockInvoiceDetail,
				invoiceStore:       mockInvoice,
			},
			args: args{
				ctx:       context.Background(),
				invoiceId: invoiceId,
			},
			mock: func() {
				mockInvoice.
					On("FindInvoice",
						context.Background(),
						map[string]interface{}{"id": invoiceId},
						moreKeys).
					Return(&invoice, nil).
					Once()

				mockInvoiceDetail.
					On("ListInvoiceDetail",
						context.Background(),
						invoiceId).
					Return(details, nil).
					Once()
			},
			want:    &finalInvoice,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeInvoiceDetailRepo{
				invoiceDetailStore: tt.fields.invoiceDetailStore,
				invoiceStore:       tt.fields.invoiceStore,
			}
			tt.mock()

			got, err := biz.SeeInvoiceDetail(tt.args.ctx, tt.args.invoiceId)

			if tt.wantErr {
				assert.NotNil(t, err, "SeeInvoiceDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "SeeInvoiceDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "SeeInvoiceDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
