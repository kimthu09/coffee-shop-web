package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockCreateToppingStore struct {
	mock.Mock
}

func (m *mockCreateToppingStore) CreateTopping(
	ctx context.Context,
	data *productmodel.ToppingCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreateRecipeStore struct {
	mock.Mock
}

func (m *mockCreateRecipeStore) CreateRecipe(
	ctx context.Context,
	data *recipemodel.RecipeCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreateListRecipeDetailStore struct {
	mock.Mock
}

func (m *mockCreateListRecipeDetailStore) CreateListRecipeDetail(
	ctx context.Context,
	data []recipedetailmodel.RecipeDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateToppingRepo(t *testing.T) {
	type args struct {
		toppingStore      CreateToppingStore
		recipeStore       CreateRecipeStore
		recipeDetailStore CreateListRecipeDetailStore
	}

	mockToppingStore := new(mockCreateToppingStore)
	mockRecipeStore := new(mockCreateRecipeStore)
	mockRecipeDetailStore := new(mockCreateListRecipeDetailStore)

	tests := []struct {
		name string
		args args
		want *createToppingRepo
	}{
		{
			name: "Create object has type CreateToppingRepo",
			args: args{
				toppingStore:      mockToppingStore,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			want: &createToppingRepo{
				toppingStore:      mockToppingStore,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateToppingRepo(tt.args.toppingStore, tt.args.recipeStore, tt.args.recipeDetailStore)

			assert.Equal(t, tt.want, got, "NewCreateToppingRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createToppingRepo_StoreTopping(t *testing.T) {
	type fields struct {
		toppingStore      CreateToppingStore
		recipeStore       CreateRecipeStore
		recipeDetailStore CreateListRecipeDetailStore
	}
	type args struct {
		ctx  context.Context
		data *productmodel.ToppingCreate
	}

	mockToppingStore := new(mockCreateToppingStore)
	mockRecipeStore := new(mockCreateRecipeStore)
	mockRecipeDetailStore := new(mockCreateListRecipeDetailStore)

	fakeToppingCreate := &productmodel.ToppingCreate{
		ProductCreate: &productmodel.ProductCreate{
			Name:         "Topping1",
			Description:  "Description1",
			CookingGuide: "",
		},
		Cost:  100,
		Price: 150,
		Recipe: &recipemodel.RecipeCreate{
			Details: []recipedetailmodel.RecipeDetailCreate{
				{
					IngredientId: "Ingredient1",
					AmountNeed:   2,
				},
				{
					IngredientId: "Ingredient2",
					AmountNeed:   3,
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
			name: "StoreTopping failed due to recipe store error",
			fields: fields{
				toppingStore:      mockToppingStore,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeToppingCreate,
			},
			mock: func() {
				mockRecipeStore.
					On("CreateRecipe", context.Background(), fakeToppingCreate.Recipe).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "StoreTopping failed due to recipe detail store error",
			fields: fields{
				toppingStore:      mockToppingStore,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeToppingCreate,
			},
			mock: func() {
				mockRecipeStore.
					On("CreateRecipe", context.Background(), fakeToppingCreate.Recipe).
					Return(nil).
					Once()

				mockRecipeDetailStore.
					On("CreateListRecipeDetail", context.Background(), fakeToppingCreate.Recipe.Details).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "StoreTopping failed due to topping store error",
			fields: fields{
				toppingStore:      mockToppingStore,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeToppingCreate,
			},
			mock: func() {
				mockRecipeStore.
					On("CreateRecipe", context.Background(), fakeToppingCreate.Recipe).
					Return(nil).
					Once()

				mockRecipeDetailStore.
					On("CreateListRecipeDetail", context.Background(), fakeToppingCreate.Recipe.Details).
					Return(nil).
					Once()

				mockToppingStore.
					On("CreateTopping", context.Background(), fakeToppingCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "StoreTopping successfully",
			fields: fields{
				toppingStore:      mockToppingStore,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeToppingCreate,
			},
			mock: func() {
				mockRecipeStore.
					On("CreateRecipe", context.Background(), fakeToppingCreate.Recipe).
					Return(nil).
					Once()

				mockRecipeDetailStore.
					On("CreateListRecipeDetail", context.Background(), fakeToppingCreate.Recipe.Details).
					Return(nil).
					Once()

				mockToppingStore.
					On("CreateTopping", context.Background(), fakeToppingCreate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createToppingRepo{
				toppingStore:      tt.fields.toppingStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := repo.StoreTopping(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "StoreTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "StoreTopping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
