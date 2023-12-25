package supplierbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel/filter"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListSupplierRepo struct {
	mock.Mock
}

func (m *mockListSupplierRepo) ListSupplier(
	ctx context.Context,
	filter *filter.Filter,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	args := m.Called(ctx, filter, paging)
	return args.Get(0).([]suppliermodel.Supplier), args.Error(1)
}

func TestNewListSupplierRepo(t *testing.T) {
	type args struct {
		repo      ListSupplierRepo
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockRepo := new(mockListSupplierRepo)

	tests := []struct {
		name string
		args args
		want *listSupplierBiz
	}{
		{
			name: "Create object has type ListSupplierBiz",
			args: args{
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &listSupplierBiz{
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListSupplierRepo(tt.args.repo, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListSupplierRepo() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listSupplierBiz_ListSupplier(t *testing.T) {
	type fields struct {
		repo      ListSupplierRepo
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *filter.Filter
		paging *common.Paging
	}

	mockRepo := new(mockListSupplierRepo)
	mockRequest := new(mockRequester)

	paging := common.Paging{
		Page: 1,
	}
	filterSupplier := filter.Filter{
		SearchKey: "",
		MinDebt:   nil,
		MaxDebt:   nil,
	}
	listSuppliers := make([]suppliermodel.Supplier, 0)
	var emptyListSuppliers []suppliermodel.Supplier
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []suppliermodel.Supplier
		wantErr bool
	}{
		{
			name: "List supplier failed because user is not allowed",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filterSupplier,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listSuppliers,
			wantErr: true,
		},
		{
			name: "List supplier failed because can not get list from database",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filterSupplier,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListSupplier",
						context.Background(),
						&filterSupplier,
						&paging).
					Return(emptyListSuppliers, mockErr).
					Once()
			},
			want:    listSuppliers,
			wantErr: true,
		},
		{
			name: "List supplier successfully",
			fields: fields{
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filterSupplier,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.SupplierViewFeatureCode).
					Return(true).
					Once()

				mockRepo.
					On(
						"ListSupplier",
						context.Background(),
						&filterSupplier,
						&paging).
					Return(listSuppliers, nil).
					Once()
			},
			want:    listSuppliers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listSupplierBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListSupplier(
				tt.args.ctx,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListSupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListSupplier() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListIngredientDetail() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
