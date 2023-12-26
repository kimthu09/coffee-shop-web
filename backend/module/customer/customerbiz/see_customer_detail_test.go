package customerbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/customer/customermodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockSeeCustomerRepo struct {
	mock.Mock
}

func (m *mockSeeCustomerRepo) SeeCustomerDetail(
	ctx context.Context,
	customerId string) (*customermodel.Customer, error) {
	args := m.Called(ctx, customerId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*customermodel.Customer), args.Error(1)
}

func TestNewSeeCustomerDetailBiz(t *testing.T) {
	type args struct {
		repo      SeeCustomerRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeCustomerRepo)

	tests := []struct {
		name string
		args args
		want *seeCustomerDetailBiz
	}{
		{
			name: "Create object has type SeeCustomerDetailBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &seeCustomerDetailBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSeeCustomerDetailBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSeeCustomerDetailBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_seeCustomerDetailBiz_SeeCustomerDetail(t *testing.T) {
	type fields struct {
		repo      SeeCustomerRepo
		requester middleware.Requester
	}
	type args struct {
		ctx        context.Context
		customerId string
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockSeeCustomerRepo)

	customerId := mock.Anything
	customer := customermodel.Customer{
		Id:       customerId,
		Name:     mock.Anything,
		Email:    mock.Anything,
		Phone:    mock.Anything,
		Point:    0,
		Invoices: nil,
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *customermodel.Customer
		wantErr bool
	}{
		{
			name: "See customer detail failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
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
			name: "See customer detail failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeCustomerDetail",
						context.Background(),
						customerId).
					Return(nil, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "See customer detail failed because can not get data from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:        context.Background(),
				customerId: customerId,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"SeeCustomerDetail",
						context.Background(),
						customerId).
					Return(&customer, nil).
					Once()
			},
			want:    &customer,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &seeCustomerDetailBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.SeeCustomerDetail(tt.args.ctx, tt.args.customerId)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"SeeCustomerDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"SeeCustomerDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"SeeCustomerDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
