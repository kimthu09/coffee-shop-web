package productbiz

import (
	"coffee_shop_management_backend/common"
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

type mockUpdateToppingRepo struct {
	mock.Mock
}

func (m *mockUpdateToppingRepo) FindTopping(
	ctx context.Context,
	id string) (*productmodel.Topping, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Topping), args.Error(1)
}

func (m *mockUpdateToppingRepo) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *mockUpdateToppingRepo) UpdateRecipeDetailsOfRecipe(
	ctx context.Context,
	recipeId string,
	deletedRecipeDetails []recipedetailmodel.RecipeDetail,
	updatedRecipeDetails []recipedetailmodel.RecipeDetailUpdate,
	createdRecipeDetails []recipedetailmodel.RecipeDetailCreate) error {
	args := m.Called(ctx, recipeId, deletedRecipeDetails, updatedRecipeDetails, createdRecipeDetails)
	return args.Error(0)
}

func (m *mockUpdateToppingRepo) FindRecipeDetails(
	ctx context.Context,
	recipeId string) ([]recipedetailmodel.RecipeDetail, error) {
	args := m.Called(ctx, recipeId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]recipedetailmodel.RecipeDetail), args.Error(1)
}

func TestNewUpdateToppingBiz(t *testing.T) {
	type args struct {
		repo      UpdateToppingRepo
		requester middleware.Requester
	}

	repo := new(mockUpdateToppingRepo)
	requester := new(mockRequester)

	tests := []struct {
		name string
		args args
		want *updateToppingBiz
	}{
		{
			name: "Create object has type UpdateToppingBiz",
			args: args{
				repo:      repo,
				requester: requester,
			},
			want: &updateToppingBiz{
				repo:      repo,
				requester: requester,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateToppingBiz(
				tt.args.repo,
				tt.args.requester,
			)
			assert.Equal(t, tt.want, got, "NewUpdateToppingBiz() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateToppingBiz_UpdateTopping(t *testing.T) {
	type fields struct {
		repo      UpdateToppingRepo
		requester middleware.Requester
	}
	type args struct {
		ctx  context.Context
		id   string
		data *productmodel.ToppingUpdateInfo
	}

	repo := new(mockUpdateToppingRepo)
	requester := new(mockRequester)

	toppingId := "Topping001"
	recipeId := "Recipe001"

	name := "updated name"
	description := "updated description"
	cookingGuide := "updated cooking guide"
	cost := 100000
	invalidCost := -1000
	price := 120000
	recipe := recipemodel.RecipeUpdate{Details: []recipedetailmodel.RecipeDetailUpdate{
		{
			IngredientId: "Ing001",
			AmountNeed:   100,
		},
		{
			IngredientId: "Ing003",
			AmountNeed:   100,
		},
	}}
	currentRecipeDetails := []recipedetailmodel.RecipeDetail{
		{
			RecipeId:     recipeId,
			IngredientId: "Ing001",
			AmountNeed:   150,
		},
		{
			RecipeId:     recipeId,
			IngredientId: "Ing004",
			AmountNeed:   300,
		},
	}
	deletedRecipeDetails := []recipedetailmodel.RecipeDetail{
		{
			RecipeId:     recipeId,
			IngredientId: "Ing004",
			AmountNeed:   300,
		},
	}
	updatedRecipeDetails := []recipedetailmodel.RecipeDetailUpdate{
		{
			IngredientId: "Ing001",
			AmountNeed:   100,
		},
	}
	createdRecipeDetails := []recipedetailmodel.RecipeDetailCreate{
		{
			RecipeId:     recipeId,
			IngredientId: "Ing003",
			AmountNeed:   100,
		},
	}
	toppingUpdate := productmodel.ToppingUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &name,
			Description:  &description,
			CookingGuide: &cookingGuide,
		},
		Cost:   &cost,
		Price:  &price,
		Recipe: &recipe,
	}
	toppingUpdateInvalid := productmodel.ToppingUpdateInfo{
		Cost: &invalidCost,
	}
	topping := productmodel.Topping{
		Product: &productmodel.Product{
			IsActive: true,
		},
		RecipeId: recipeId,
	}
	inactiveTopping := productmodel.Topping{
		Product: &productmodel.Product{
			IsActive: false,
		},
		RecipeId: recipeId,
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
			name: "Update topping failed because user is not allowed",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(false).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping failed because data is not valid",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdateInvalid,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping failed because can not find topping",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
				repo.
					On("FindTopping",
						context.Background(),
						toppingId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping failed because topping is inactive",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
				repo.
					On("FindTopping",
						context.Background(),
						toppingId).
					Return(&inactiveTopping, nil).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping failed because can not save general info of topping into database",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
				repo.
					On("FindTopping",
						context.Background(),
						toppingId).
					Return(&topping, nil).
					Once()

				repo.
					On("UpdateTopping",
						context.Background(),
						toppingId,
						&toppingUpdate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping failed because can not get current recipe's details",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
				repo.
					On("FindTopping",
						context.Background(),
						toppingId).
					Return(&topping, nil).
					Once()

				repo.
					On("UpdateTopping",
						context.Background(),
						toppingId,
						&toppingUpdate).
					Return(nil).
					Once()

				repo.
					On("FindRecipeDetails",
						context.Background(),
						recipeId).
					Return(nil, mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping failed because can not get update recipe's details",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
				repo.
					On("FindTopping",
						context.Background(),
						toppingId).
					Return(&topping, nil).
					Once()

				repo.
					On("UpdateTopping",
						context.Background(),
						toppingId,
						&toppingUpdate).
					Return(nil).
					Once()

				repo.
					On("FindRecipeDetails",
						context.Background(),
						recipeId).
					Return(currentRecipeDetails, nil).
					Once()

				repo.
					On("UpdateRecipeDetailsOfRecipe",
						context.Background(),
						recipeId,
						deletedRecipeDetails,
						updatedRecipeDetails,
						createdRecipeDetails).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping successfully",
			fields: fields{
				repo:      repo,
				requester: requester,
			},
			args: args{
				ctx:  context.Background(),
				id:   toppingId,
				data: &toppingUpdate,
			},
			mock: func() {
				requester.
					On("IsHasFeature",
						common.ToppingUpdateInfoFeatureCode).
					Return(true).
					Once()
				repo.
					On("FindTopping",
						context.Background(),
						toppingId).
					Return(&topping, nil).
					Once()

				repo.
					On("UpdateTopping",
						context.Background(),
						toppingId,
						&toppingUpdate).
					Return(nil).
					Once()

				repo.
					On("FindRecipeDetails",
						context.Background(),
						recipeId).
					Return(currentRecipeDetails, nil).
					Once()

				repo.
					On("UpdateRecipeDetailsOfRecipe",
						context.Background(),
						recipeId,
						deletedRecipeDetails,
						updatedRecipeDetails,
						createdRecipeDetails).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateToppingBiz{
				repo:      tt.fields.repo,
				requester: tt.fields.requester,
			}

			tt.mock()

			err := biz.UpdateTopping(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateTopping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
