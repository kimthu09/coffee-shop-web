package ingredientbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListIngredientStore struct {
	mock.Mock
}

func (m *mockListIngredientStore) ListIngredient(
	ctx context.Context,
	filter *ingredientmodel.Filter,
	propertiesContainSearchKey []string,
	paging *common.Paging) ([]ingredientmodel.Ingredient, error) {
	args := m.Called(ctx, filter, propertiesContainSearchKey, paging)
	return args.Get(0).([]ingredientmodel.Ingredient), args.Error(1)
}

func TestNewListIngredientBiz(t *testing.T) {
	type args struct {
		store     ListIngredientStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListIngredientStore)

	tests := []struct {
		name string
		args args
		want *listIngredientBiz
	}{
		{
			name: "Create object has type ListIngredientBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listIngredientBiz{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListIngredientBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListIngredientBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listIngredientBiz_ListImportNote(t *testing.T) {
	type fields struct {
		store     ListIngredientStore
		requester middleware.Requester
	}
	type args struct {
		ctx    context.Context
		filter *ingredientmodel.Filter
		paging *common.Paging
	}

	mockStore := new(mockListIngredientStore)
	mockRequest := new(mockRequester)
	filter := ingredientmodel.Filter{
		SearchKey:   "",
		MinPrice:    nil,
		MaxPrice:    nil,
		MinAmount:   nil,
		MaxAmount:   nil,
		MeasureType: "",
	}
	paging := common.Paging{
		Page: 1,
	}
	properties := []string{"id", "name"}
	listIngredients := make([]ingredientmodel.Ingredient, 0)
	var emptyListIngredients []ingredientmodel.Ingredient
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []ingredientmodel.Ingredient
		wantErr bool
	}{
		{
			name: "List ingredient failed because user is not allowed",
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
					On("IsHasFeature", common.IngredientViewFeatureCode).
					Return(false).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List ingredient failed because can not get data from database",
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
					On("IsHasFeature", common.IngredientViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListIngredient",
						context.Background(),
						&filter,
						properties,
						&paging,
					).
					Return(emptyListIngredients, mockErr).
					Once()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List ingredient successfully",
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
					On("IsHasFeature", common.IngredientViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListIngredient",
						context.Background(),
						&filter,
						properties,
						&paging,
					).
					Return(listIngredients, nil).
					Once()
			},
			want:    listIngredients,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listIngredientBiz{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListIngredient(
				tt.args.ctx,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListIngredient() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListIngredient() error = %v, wantErr %v",
					err,
					tt.wantErr)
				assert.Equal(
					t,
					tt.want,
					got,
					"ListIngredient() want = %v, got %v",
					tt.want,
					got)
			}
		})
	}
}
