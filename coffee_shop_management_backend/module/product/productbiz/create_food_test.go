package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateFoodRepo struct {
	mock.Mock
}

func (m *mockCreateFoodRepo) CreateFood(ctx context.Context, data *productmodel.FoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateFoodRepo) HandleCategoryFood(ctx context.Context, foodId string, data *productmodel.FoodCreate) error {
	args := m.Called(ctx, foodId, data)
	return args.Error(0)
}

func (m *mockCreateFoodRepo) HandleSizeFood(ctx context.Context, data *productmodel.FoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockIdGenerator struct {
	mock.Mock
}

func (m *mockIdGenerator) GenerateId() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *mockIdGenerator) IdProcess(id *string) (*string, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*string), args.Error(1)
}

func TestNewCreateFoodBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      CreateFoodRepo
		requester middleware.Requester
	}

	mockRepo := new(mockCreateFoodRepo)
	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *createFoodBiz
	}{
		{
			name: "Create object has type CreateFoodBiz",
			args: args{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			want: &createFoodBiz{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateFoodBiz(tt.args.gen, tt.args.repo, tt.args.requester)
			assert.Equal(t, tt.want, got, "NewCreateFoodBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createFoodBiz_CreateFood(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      CreateFoodRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		data *productmodel.FoodCreate
	}

	mockRepo := new(mockCreateFoodRepo)
	mockGenerator := new(mockIdGenerator)
	mockRequest := new(mockRequester)

	foodId := "Food001"
	sizeId := "Size001"
	recipeId := "Recipe001"
	categories := []string{"Category001"}
	sizes := []sizefoodmodel.SizeFoodCreate{
		{
			Name:  "S",
			Cost:  10000,
			Price: 8000,
			Recipe: &recipemodel.RecipeCreate{
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						IngredientId: "Ing001",
						AmountNeed:   10,
					},
				},
			},
		},
	}
	validData := &productmodel.FoodCreate{
		ProductCreate: &productmodel.ProductCreate{
			Id:           nil,
			Name:         "Food 1",
			Description:  "",
			CookingGuide: "",
		},
		Categories: categories,
		Sizes:      sizes,
	}

	finalParam := &productmodel.FoodCreate{
		ProductCreate: &productmodel.ProductCreate{
			Id:           &foodId,
			Name:         "Food 1",
			Description:  "",
			CookingGuide: "",
		},
		Categories: categories,
		Sizes: []sizefoodmodel.SizeFoodCreate{
			{
				FoodId:   foodId,
				SizeId:   sizeId,
				Name:     "S",
				Cost:     10000,
				Price:    8000,
				RecipeId: recipeId,
				Recipe: &recipemodel.RecipeCreate{
					Id: recipeId,
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							RecipeId:     "Recipe001",
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
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
			name: "Create food failed because user is not allowed",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food failed because data is invalid",
			fields: fields{
				gen:       mockGenerator,
				repo:      mockRepo,
				requester: mockRequest,
			},
			args: args{
				ctx: context.Background(),
				data: &productmodel.FoodCreate{
					ProductCreate: &productmodel.ProductCreate{
						Name: "",
					},
				},
			},
			mock: func() {
				mockRequest.
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food failed because can not handle food id",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
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
			name: "Create food failed because can not generate size id",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&foodId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food failed because can not generate recipe id",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&foodId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(sizeId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food failed because can not create food",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&foodId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(sizeId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(recipeId, nil).
					Once()

				mockRepo.
					On("CreateFood", context.Background(), validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food failed because can not handle category food",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&foodId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(sizeId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(recipeId, nil).
					Once()

				mockRepo.
					On("CreateFood", context.Background(), validData).
					Return(nil).
					Once()

				mockRepo.
					On("HandleCategoryFood", context.Background(), foodId, validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food failed because can not handle size food",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&foodId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(sizeId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(recipeId, nil).
					Once()

				mockRepo.
					On("CreateFood", context.Background(), validData).
					Return(nil).
					Once()

				mockRepo.
					On("HandleCategoryFood", context.Background(), foodId, validData).
					Return(nil).
					Once()

				mockRepo.
					On("HandleSizeFood", context.Background(), validData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create food successfully",
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
					On("IsHasFeature", common.FoodCreateFeatureCode).
					Return(true).
					Once()

				mockGenerator.
					On("IdProcess", validData.Id).
					Return(&foodId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(sizeId, nil).
					Once()

				mockGenerator.
					On("GenerateId").
					Return(recipeId, nil).
					Once()

				mockRepo.
					On("CreateFood", context.Background(), validData).
					Return(nil).
					Once()

				mockRepo.
					On("HandleCategoryFood", context.Background(), foodId, validData).
					Return(nil).
					Once()

				mockRepo.
					On("HandleSizeFood", context.Background(), validData).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &createFoodBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.CreateFood(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.args.data, finalParam, "Param = %, want %v", tt.args.data, finalParam)
			}
		})
	}
}
