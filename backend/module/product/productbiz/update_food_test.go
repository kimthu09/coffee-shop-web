package productbiz

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/component/generator"
	"coffee_shop_management_backend/middleware"
	"coffee_shop_management_backend/module/category/categorymodel"
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

type mockUpdateFoodRepo struct {
	mock.Mock
}

func (m *mockUpdateFoodRepo) FindFood(
	ctx context.Context,
	id string) (*productmodel.Food, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Food), args.Error(1)
}

func (m *mockUpdateFoodRepo) FindCategories(
	ctx context.Context,
	foodId string) ([]categorymodel.SimpleCategoryWithId, error) {
	args := m.Called(ctx, foodId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]categorymodel.SimpleCategoryWithId), args.Error(1)
}

func (m *mockUpdateFoodRepo) HandleCategory(
	ctx context.Context,
	foodId string,
	deletedCategoryFood []categorymodel.SimpleCategoryWithId,
	createdCategoryFood []categorymodel.SimpleCategoryWithId) error {
	args := m.Called(ctx, foodId, deletedCategoryFood, createdCategoryFood)
	return args.Error(0)
}

func (m *mockUpdateFoodRepo) FindSizeFoods(
	ctx context.Context,
	foodId string) ([]sizefoodmodel.SizeFood, error) {
	args := m.Called(ctx, foodId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]sizefoodmodel.SizeFood), args.Error(1)
}

func (m *mockUpdateFoodRepo) FindRecipeDetails(
	ctx context.Context,
	recipeId string) ([]recipedetailmodel.RecipeDetail, error) {
	args := m.Called(ctx, recipeId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]recipedetailmodel.RecipeDetail), args.Error(1)
}

func (m *mockUpdateFoodRepo) HandleSizeFoods(
	ctx context.Context,
	foodId string,
	deletedSizeFood []sizefoodmodel.SizeFood,
	updatedSizeFood []sizefoodmodel.SizeFoodUpdate,
	mapDeletedRecipeDetails map[string][]recipedetailmodel.RecipeDetail,
	mapUpdatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailUpdate,
	mapCreatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailCreate,
	createdSizeFood []sizefoodmodel.SizeFoodCreate) error {
	args := m.Called(ctx, foodId, deletedSizeFood, updatedSizeFood, mapDeletedRecipeDetails, mapUpdatedRecipeDetails, mapCreatedRecipeDetails, createdSizeFood)
	return args.Error(0)
}

func (m *mockUpdateFoodRepo) UpdateFood(
	ctx context.Context,
	id string,
	data *productmodel.FoodUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewUpdateFoodBiz(t *testing.T) {
	type args struct {
		gen       generator.IdGenerator
		repo      UpdateFoodRepo
		requester middleware.Requester
	}

	gen := new(mockIdGenerator)
	repo := new(mockUpdateFoodRepo)
	requester := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateFoodBiz
	}{
		{
			name: "Create object has type UpdateFoodBiz",
			args: args{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			want: &updateFoodBiz{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateFoodBiz(
				tt.args.gen,
				tt.args.repo,
				tt.args.requester,
			)
			assert.Equal(t, tt.want, got, "NewUpdateFoodBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateFoodBiz_UpdateFood(t *testing.T) {
	type fields struct {
		gen       generator.IdGenerator
		repo      UpdateFoodRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *productmodel.FoodUpdateInfo
	}

	gen := new(mockIdGenerator)
	repo := new(mockUpdateFoodRepo)
	requester := new(mockRequester)

	foodId := "Food001"
	name := "updated name"
	description := "updated description"
	cookingGuide := "updated cooking guide"
	sizeId := "Size001"
	sizeName := "updated size name"
	sizeCost := 10000
	sizePrice := 12000
	recipeId := "Recipe001"
	recipe := recipemodel.RecipeUpdate{Details: []recipedetailmodel.RecipeDetailUpdate{
		{
			IngredientId: "Ing001",
			AmountNeed:   100,
		},
		{
			IngredientId: "Ing002",
			AmountNeed:   100,
		},
	}}

	currentRecipeDetails := []recipedetailmodel.RecipeDetail{
		{
			RecipeId:     recipeId,
			IngredientId: "Ing002",
			AmountNeed:   150,
		},
		{
			RecipeId:     recipeId,
			IngredientId: "Ing003",
			AmountNeed:   300,
		},
	}
	deletedSizeId := "SizeDeleted"
	currentSizeFoods := []sizefoodmodel.SizeFood{
		{
			FoodId:   foodId,
			SizeId:   sizeId,
			RecipeId: recipeId,
			Recipe: recipemodel.Recipe{
				Id:      recipeId,
				Details: currentRecipeDetails,
			},
		},
		{
			FoodId: foodId,
			SizeId: deletedSizeId,
		},
	}

	currentCategories := []categorymodel.SimpleCategoryWithId{
		{
			"Cat002",
		},
		{
			"Cat003",
		},
	}
	deletedCategories := []categorymodel.SimpleCategoryWithId{
		{
			"Cat003",
		},
	}
	createdCategories := []categorymodel.SimpleCategoryWithId{
		{
			"Cat001",
		},
	}
	sizeFoodUpdate := []sizefoodmodel.SizeFoodUpdate{
		{
			SizeId:   &sizeId,
			Name:     &sizeName,
			Cost:     &sizeCost,
			Price:    &sizePrice,
			RecipeId: &recipeId,
			Recipe:   &recipe,
		},
		{
			SizeId:   nil,
			Name:     &sizeName,
			Cost:     &sizeCost,
			Price:    &sizePrice,
			RecipeId: nil,
			Recipe:   &recipe,
		},
	}

	foodUpdate := productmodel.FoodUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &name,
			Description:  &description,
			CookingGuide: &cookingGuide,
		},
		Categories: &[]string{"Cat001", "Cat002"},
		Sizes:      &sizeFoodUpdate,
	}

	deletedSizes := []sizefoodmodel.SizeFood{
		{
			FoodId: foodId,
			SizeId: deletedSizeId,
		},
	}
	updatedSizes := []sizefoodmodel.SizeFoodUpdate{
		{
			SizeId:   &sizeId,
			Name:     &sizeName,
			Cost:     &sizeCost,
			Price:    &sizePrice,
			RecipeId: &recipeId,
			Recipe:   &recipe,
		},
	}
	mapDeletedRecipeDetails := map[string][]recipedetailmodel.RecipeDetail{
		recipeId: {
			{
				RecipeId:     recipeId,
				IngredientId: "Ing003",
				AmountNeed:   300,
			},
		},
	}
	mapUpdatedRecipeDetails := map[string][]recipedetailmodel.RecipeDetailUpdate{
		recipeId: {
			{
				IngredientId: "Ing002",
				AmountNeed:   100,
			},
		},
	}
	mapCreatedRecipeDetails := map[string][]recipedetailmodel.RecipeDetailCreate{
		recipeId: {
			{
				RecipeId:     recipeId,
				IngredientId: "Ing001",
				AmountNeed:   100,
			},
		},
	}
	newSizeId := "SizeCreate"
	newRecipeId := "RecipeCreate"
	createdSizes := []sizefoodmodel.SizeFoodCreate{
		{
			FoodId:   foodId,
			SizeId:   newSizeId,
			Name:     sizeName,
			Cost:     sizeCost,
			Price:    sizePrice,
			RecipeId: newRecipeId,
			Recipe: &recipemodel.RecipeCreate{
				Id: newRecipeId,
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						RecipeId:     newRecipeId,
						IngredientId: "Ing001",
						AmountNeed:   100,
					},
					{
						RecipeId:     newRecipeId,
						IngredientId: "Ing002",
						AmountNeed:   100,
					},
				},
			},
		},
	}

	emptyString := ""
	invalidSizeUpdate := "NewSize"

	sizeFoodUpdateInvalidSizeCreate := []sizefoodmodel.SizeFoodUpdate{
		{
			SizeId:   nil,
			Name:     &sizeName,
			Cost:     nil,
			Price:    &sizePrice,
			RecipeId: nil,
			Recipe:   &recipe,
		},
	}
	foodUpdateHasInvalidSizeCreate := productmodel.FoodUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &name,
			Description:  &description,
			CookingGuide: &cookingGuide,
		},
		Categories: &[]string{"Cat001", "Cat002"},
		Sizes:      &sizeFoodUpdateInvalidSizeCreate,
	}

	sizeFoodUpdateInvalidSizeUpdate := []sizefoodmodel.SizeFoodUpdate{
		{
			SizeId:   &invalidSizeUpdate,
			Name:     &sizeName,
			Cost:     &sizeCost,
			Price:    &sizePrice,
			RecipeId: nil,
			Recipe:   &recipe,
		},
	}
	foodUpdateHasInvalidSizeUpdate := productmodel.FoodUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &name,
			Description:  &description,
			CookingGuide: &cookingGuide,
		},
		Categories: &[]string{"Cat001", "Cat002"},
		Sizes:      &sizeFoodUpdateInvalidSizeUpdate,
	}

	food := productmodel.Food{
		Product: &productmodel.Product{
			Id:       foodId,
			IsActive: true,
		},
	}
	inactiveFood := productmodel.Food{
		Product: &productmodel.Product{
			Id:       foodId,
			IsActive: false,
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
			name: "Update food failed because user is not allowed",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because data is not valid",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx: context.Background(),
				id:  foodId,
				data: &productmodel.FoodUpdateInfo{
					ProductUpdateInfo: &productmodel.ProductUpdateInfo{
						Name: &emptyString,
					},
				},
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not find food",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because food is inactive",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&inactiveFood, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because food is inactive",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&inactiveFood, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not get current categories",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not save category into database",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not find current size food",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(nil, mockErr).
					Once()

				gen.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because the size need to update is not valid",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdateHasInvalidSizeUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not generate size id",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not generate recipe id",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newSizeId, nil).
					Once()

				gen.
					On("GenerateId").
					Return("", mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because the size need to create is not valid",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdateHasInvalidSizeCreate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newSizeId, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newRecipeId, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not find current recipe's details to update size food",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newSizeId, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newRecipeId, nil).
					Once()

				repo.
					On(
						"FindRecipeDetails",
						context.Background(),
						recipeId,
					).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not save size foods into database",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newSizeId, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newRecipeId, nil).
					Once()

				repo.
					On(
						"FindRecipeDetails",
						context.Background(),
						recipeId,
					).
					Return(currentRecipeDetails, nil).
					Once()

				repo.
					On(
						"HandleSizeFoods",
						context.Background(),
						foodId,
						deletedSizes,
						updatedSizes,
						mapDeletedRecipeDetails,
						mapUpdatedRecipeDetails,
						mapCreatedRecipeDetails,
						createdSizes,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food failed because can not save general info into database",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newSizeId, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newRecipeId, nil).
					Once()

				repo.
					On(
						"FindRecipeDetails",
						context.Background(),
						recipeId,
					).
					Return(currentRecipeDetails, nil).
					Once()

				repo.
					On(
						"HandleSizeFoods",
						context.Background(),
						foodId,
						deletedSizes,
						updatedSizes,
						mapDeletedRecipeDetails,
						mapUpdatedRecipeDetails,
						mapCreatedRecipeDetails,
						createdSizes,
					).
					Return(nil).
					Once()

				repo.
					On(
						"UpdateFood",
						context.Background(),
						foodId,
						&foodUpdate,
					).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food successfully",
			fields: fields{
				gen:       gen,
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   foodId,
				data: &foodUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.FoodUpdateInfoFeatureCode).
					Return(true).
					Once()

				repo.
					On("FindFood",
						context.Background(),
						foodId).
					Return(&food, nil).
					Once()

				repo.
					On("FindCategories",
						context.Background(),
						foodId).
					Return(currentCategories, nil).
					Once()

				repo.
					On("HandleCategory",
						context.Background(),
						foodId,
						deletedCategories,
						createdCategories).
					Return(nil).
					Once()

				repo.
					On("FindSizeFoods",
						context.Background(),
						foodId).
					Return(currentSizeFoods, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newSizeId, nil).
					Once()

				gen.
					On("GenerateId").
					Return(newRecipeId, nil).
					Once()

				repo.
					On(
						"FindRecipeDetails",
						context.Background(),
						recipeId,
					).
					Return(currentRecipeDetails, nil).
					Once()

				repo.
					On(
						"HandleSizeFoods",
						context.Background(),
						foodId,
						deletedSizes,
						updatedSizes,
						mapDeletedRecipeDetails,
						mapUpdatedRecipeDetails,
						mapCreatedRecipeDetails,
						createdSizes,
					).
					Return(nil).
					Once()

				repo.
					On(
						"UpdateFood",
						context.Background(),
						foodId,
						&foodUpdate,
					).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateFoodBiz{
				gen:       tt.fields.gen,
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}
			tt.mock()

			err := biz.UpdateFood(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
