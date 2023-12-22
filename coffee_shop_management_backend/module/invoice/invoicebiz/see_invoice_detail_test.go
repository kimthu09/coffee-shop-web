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

type mocSeeInvoiceDetailRepo struct {
	mock.Mock
}

func (m *mocSeeInvoiceDetailRepo) SeeInvoiceDetail(
	ctx context.Context,
	invoiceId string) (*invoicemodel.Invoice, error) {
	args := m.Called(ctx, invoiceId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*invoicemodel.Invoice), args.Error(1)
}

type mockRequester struct {
	mock.Mock
}

func (m *mockRequester) GetUserId() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetEmail() string {
	args := m.Called()
	return args.String(0)
}
func (m *mockRequester) GetRoleId() string {
	args := m.Called()
	return args.Get(0).(string)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewSeeInvoiceDetailBiz(t *testing.T) {
	type args struct {
		repo      SeeInvoiceDetailRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mocSeeInvoiceDetailRepo)

	tests := []struct {
		name string
		args args
		want *seeInvoiceDetailBiz
	}{
		{
			name: "Create object has type SeeInvoiceDetailBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeInvoiceDetailBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeInvoiceDetailBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeInvoiceDetailBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeInvoiceDetailBiz_SeeInvoiceDetail(t *testing.T) {
	type fields struct {
		repo      SeeInvoiceDetailRepo
		requester middleware.Requester
	}
	type args struct {
		ctx       context.Context
		invoiceId string
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mocSeeInvoiceDetailRepo)

	invoice := invoicemodel.Invoice{
		Id: "Invoice001",
	}
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
			name: "See invoice detail failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:       context.Background(),
				invoiceId: "Invoice001",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceViewFeatureCode).
					Return(false).
					Once()
			},
			want:    &invoice,
			wantErr: true,
		},
		{
			name: "See invoice detail failed because can not get supplier from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:       context.Background(),
				invoiceId: "Invoice001",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeInvoiceDetail",
						context.Background(),
						"Invoice001").
					Return(nil, mockErr).
					Once()
			},
			want:    &invoice,
			wantErr: true,
		},
		{
			name: "See invoice detail successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:       context.Background(),
				invoiceId: "Invoice001",
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.InvoiceViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeInvoiceDetail",
						context.Background(),
						"Invoice001").
					Return(&invoice, nil).
					Once()
			},
			want:    &invoice,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeInvoiceDetailBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeInvoiceDetail(tt.args.ctx, tt.args.invoiceId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeInvoiceDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeInvoiceDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeInvoiceDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
