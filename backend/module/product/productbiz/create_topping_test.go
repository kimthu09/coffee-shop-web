package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateToppingRepo struct {
	mock.Mock
}

func (m *mockCreateToppingRepo) StoreTopping(ctx context.Context, data *productmodel.ToppingCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateToppingBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateToppingRepo
		requester middleware.Requester
	}
	mockRepo := new(mockCreateToppingRepo)
	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createToppingBiz
	}{
		{
			name: "Create object has type CreateToppingBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createToppingBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateToppingBiz(tt.args.gen, tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewCreateToppingBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createToppingBiz_CreateTopping(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateToppingRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *productmodel.ToppingCreate
	}

	mockRepo := new(mockCreateToppingRepo)
	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)

	toppingId := "Topping001"
	recipeId := "Recipe001"
	validData := &productmodel.ToppingCreate{
		ProductCreate: &productmodel.ProductCreate{
			Id:           nil,
			Name:         "Topping 1",
			Description:  "",
			CookingGuide: "",
		},
		Cost:  0,
		Price: 0,
		Recipe: &recipemodel.RecipeCreate{
			Details: []recipedetailmodel.RecipeDetailCreate{
				{
					IngredientId: "Ing001",
					AmountNeed:   10,
				},
			},
		},
	}
	finalParam := &productmodel.ToppingCreate{
		ProductCreate: &productmodel.ProductCreate{
			Id:           &toppingId,
			Name:         "Topping 1",
			Description:  "",
			CookingGuide: "",
		},
		Cost:     0,
		Price:    0,
		RecipeId: recipeId,
		Recipe: &recipemodel.RecipeCreate{
			Id: recipeId,
			Details: []recipedetailmodel.RecipeDetailCreate{
				{
					RecipeId:     recipeId,
					IngredientId: "Ing001",
					AmountNeed:   10,
				},
			},
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create topping failed because user is not allowed",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create topping failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: &productmodel.ToppingCreate{},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create topping failed because can not handle topping id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create topping failed because can not generate recipe id",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&toppingId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create topping failed because can not create topping",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&toppingId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(recipeId, nil).
					Once()

				mockRepo.
					On("StoreTopping", context.Background(), validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create topping successfully",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx:  context.Background(),
				data: validData,
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.ToppingCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&toppingId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(recipeId, nil).
					Once()

				mockRepo.
					On("StoreTopping", context.Background(), validData).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createToppingBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateTopping(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.args.data, finalParam, "Param = %, want %v", tt.args.data, finalParam)
			}
		})
	}
}
