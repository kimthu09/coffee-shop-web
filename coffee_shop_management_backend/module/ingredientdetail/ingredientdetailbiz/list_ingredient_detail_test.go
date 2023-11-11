package ingredientdetailbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/ingredientdetail/ingredientdetailmodel"
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockListIngredientDetailStore struct {
	mock.Mock
}

func (m *mockListIngredientDetailStore) ListIngredientDetail(
	ctx context.Context,
	ingredientId string,
	filter *ingredientdetailmodel.Filter,
	paging *common.Paging) ([]ingredientdetailmodel.IngredientDetail, error) {
	args := m.Called(ctx, ingredientId, filter, paging)
	return args.Get(0).([]ingredientdetailmodel.IngredientDetail),
		args.Error(1)
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
func (m *mockRequester) GetRole() rolemodel.Role {
	args := m.Called()
	return args.Get(0).(rolemodel.Role)
}
func (m *mockRequester) IsHasFeature(featureCode string) bool {
	args := m.Called(featureCode)
	return args.Bool(0)
}

func TestNewListIngredientDetailByIdBiz(t *testing.T) {
	type args struct {
		store     ListIngredientDetailStore
		requester middleware.Requester
	}

	mockRequest := new(mockRequester)
	mockStore := new(mockListIngredientDetailStore)

	tests := []struct {
		name string
		args args
		want *listIngredientDetail
	}{
		{
			name: "Create object has type ListIngredientDetailBiz",
			args: args{
				store:     mockStore,
				requester: mockRequest,
			},
			want: &listIngredientDetail{
				store:     mockStore,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewListIngredientDetailByIdBiz(tt.args.store, tt.args.requester)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewListIngredientDetailByIdBiz() = %v, want %v",
				got,
				tt.want)
		})
	}
}

func Test_listIngredientDetail_ListIngredientDetail(t *testing.T) {
	type fields struct {
		store     ListIngredientDetailStore
		requester middleware.Requester
	}
	type args struct {
		ctx          context.Context
		ingredientId string
		filter       *ingredientdetailmodel.Filter
		paging       *common.Paging
	}

	mockStore := new(mockListIngredientDetailStore)
	mockRequest := new(mockRequester)

	paging := common.Paging{
		Page: 1,
	}
	filter := ingredientdetailmodel.Filter{IsGetEmpty: false}
	ingredientId := mock.Anything
	listIngredientDetails := make([]ingredientdetailmodel.IngredientDetail, 0)
	var emptyListIngredientDetails []ingredientdetailmodel.IngredientDetail
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []ingredientdetailmodel.IngredientDetail
		wantErr bool
	}{
		{
			name: "List ingredient detail failed because user is not allowed",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: ingredientId,
				filter:       &filter,
				paging:       &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientViewFeatureCode).
					Return(false).
					Once()
			},
			want:    listIngredientDetails,
			wantErr: true,
		},
		{
			name: "List ingredient detail failed because can not get data from database",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: ingredientId,
				filter:       &filter,
				paging:       &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListIngredientDetail",
						context.Background(),
						ingredientId,
						&filter,
						&paging,
					).
					Return(emptyListIngredientDetails, mockErr).
					Once()
			},
			want:    listIngredientDetails,
			wantErr: true,
		},
		{
			name: "List ingredient detail successfully",
			fields: fields{
				store:     mockStore,
				requester: mockRequest,
			},
			args: args{
				ctx:          context.Background(),
				ingredientId: ingredientId,
				filter:       &filter,
				paging:       &paging,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.IngredientViewFeatureCode).
					Return(true).
					Once()

				mockStore.
					On(
						"ListIngredientDetail",
						context.Background(),
						ingredientId,
						&filter,
						&paging,
					).
					Return(listIngredientDetails, nil).
					Once()
			},
			want:    listIngredientDetails,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &listIngredientDetail{
				store:     tt.fields.store,
				requester: tt.fields.requester,
			}

			tt.mock()

			got, err := biz.ListIngredientDetail(
				tt.args.ctx,
				tt.args.ingredientId,
				tt.args.filter,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(
					t,
					err,
					"ListIngredientDetail() error = %v, wantErr %v",
					err,
					tt.wantErr)
			} else {
				assert.Nil(
					t,
					err,
					"ListIngredientDetail() error = %v, wantErr %v",
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
