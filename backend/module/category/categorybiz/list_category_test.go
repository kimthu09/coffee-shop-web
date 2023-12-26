package categorybiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListCategoryStore struct {
	mock.Mock
}

func (m *mockListCategoryStore) ListCategory(
	ctx context.Context,
	filter *categorymodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]categorymodel.Category, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	return args.Get(0).([]categorymodel.Category), args.Error(1)
}

func TestNewListCategoryBiz(t *testing.T) {
	type args struct {
		store     ListCategoryStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListCategoryStore)

	tests := []struct {
		name string
		args args
		want *listCategoryBiz
	}{
		{
			name: "Create object has type ListCustomerBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listCategoryBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListCategoryBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListCategoryBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listCategoryBiz_ListCategory(t *testing.T) {
	type fields struct {
		store     ListCategoryStore
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *categorymodel.Filter
		paging *common.Paging
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListCategoryStore)

	paging := common.Paging{
		Page:  1,
		Limit: 5,
	}
	filter := categorymodel.Filter{
		SearchKey: "",
	}
	listCategories := make([]categorymodel.Category, 0)
	var emptyListCategories []categorymodel.Category
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []categorymodel.Category
		wantErr bool
	}{
		{
			name: "List category failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List category failed because can not get list from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListCategory",
						context.Background(),
						&filter,
						[]string{"id", "name"},
						&paging).
					Return(emptyListCategories, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List category successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:    context.Background(),
				filter: &filter,
				paging: &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.CategoryViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListCategory",
						context.Background(),
						&filter,
						[]string{"id", "name"},
						&paging).
					Return(listCategories, nil).
					Once()
			},
			want:    listCategories,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listCategoryBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListCategory(
				tt.args.ctx,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListCategory() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListCategory() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListCategory() = %v, want %v",
					got,
					tt.want)
			}
		})
	}
}
