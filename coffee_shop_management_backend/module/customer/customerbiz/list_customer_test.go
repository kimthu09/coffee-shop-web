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

type mockListCustomerRepo struct {
	mock.Mock
}

func (m *mockListCustomerRepo) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	paging *common.Paging) ([]customermodel.Customer, error) {
	args := m.Called(ctx, filter, paging)
	return args.Get(0).([]customermodel.Customer), args.Error(1)
}

func TestNewListCustomerBiz(t *testing.T) {
	type args struct {
		repo      ListCustomerRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListCustomerRepo)

	tests := []struct {
		name string
		args args
		want *listCustomerBiz
	}{
		{
			name: "Create object has type ListCustomerBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listCustomerBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListCustomerBiz(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListCustomerRepo() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listCustomerBiz_ListCustomer(t *testing.T) {
	type fields struct {
		repo      ListCustomerRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *customermodel.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListCustomerRepo)
	mockRequest := new(mockRequester)

	paging := common.Paging{
		Page: 1,
	}
	filter := customermodel.Filter{
		SearchKey: "",
		MinPoint:  nil,
		MaxPoint:  nil,
	}
	listCustomers := make([]customermodel.Customer, 0)
	var emptyListCustomers []customermodel.Customer
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []customermodel.Customer
		wantErr bool
	}{
		{
			name: "List customer failed because user is not allowed",
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
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listCustomers,
			wantErr: true,
		},
		{
			name: "List customer failed because can not get list from database",
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
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListCustomer",
						context.Background(),
						&filter,
						&paging).
					Return(emptyListCustomers, mockErr).
					Once()
			},
			want:    listCustomers,
			wantErr: true,
		},
		{
			name: "List customer successfully",
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
					On("IsHasFeature", common.CustomerViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListCustomer",
						context.Background(),
						&filter,
						&paging).
					Return(listCustomers, nil).
					Once()
			},
			want:    listCustomers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listCustomerBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListCustomer(
				tt.args.ctx,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListCustomer() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListCustomer() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListCustomer() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
