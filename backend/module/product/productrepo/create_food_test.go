package productrepo

import (
	"coffee_shop_management_backend/module/category/categorymodel"
	"coffee_shop_management_backend/module/categoryfood/categoryfoodmodel"
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

type mockCreateFoodStore struct {
	mock.Mock
}

func (m *mockCreateFoodStore) CreateFood(
	ctx context.Context,
	data *productmodel.FoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockCreateCategoryFoodStore struct {
	mock.Mock
}

func (m *mockCreateCategoryFoodStore) CreateCategoryFood(
	ctx context.Context,
	data *categoryfoodmodel.CategoryFoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type mockUpdateCategoryStore struct {
	mock.Mock
}

func (m *mockUpdateCategoryStore) UpdateAmountProductCategory(
	ctx context.Context,
	id string,
	data *categorymodel.CategoryUpdateAmountProduct) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

type mockCreateSizeFoodStore struct {
	mock.Mock
}

func (m *mockCreateSizeFoodStore) CreateSizeFood(
	ctx context.Context,
	data *sizefoodmodel.SizeFoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func TestNewCreateFoodRepo(t *testing.T) {
	type args struct {
		foodStore         CreateFoodStore
		categoryFoodStore CreateCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     CreateSizeFoodStore
		recipeStore       CreateRecipeStore
		recipeDetailStore CreateListRecipeDetailStore
	}

	mockFood := new(mockCreateFoodStore)
	mockCategoryFood := new(mockCreateCategoryFoodStore)
	mockCategory := new(mockUpdateCategoryStore)
	mockSizeFood := new(mockCreateSizeFoodStore)
	mockRecipeStore := new(mockCreateRecipeStore)
	mockRecipeDetailStore := new(mockCreateListRecipeDetailStore)

	tests := []struct {
		name string
		args args
		want *createFoodRepo
	}{
		{
			name: "Create object has type CreateFoodRepo",
			args: args{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			want: &createFoodRepo{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCreateFoodRepo(tt.args.foodStore, tt.args.categoryFoodStore, tt.args.categoryStore, tt.args.sizeFoodStore, tt.args.recipeStore, tt.args.recipeDetailStore)

			assert.Equal(t, tt.want, got, "NewCreateFoodRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_createFoodRepo_CreateFood(t *testing.T) {
	type fields struct {
		foodStore         CreateFoodStore
		categoryFoodStore CreateCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     CreateSizeFoodStore
		recipeStore       CreateRecipeStore
		recipeDetailStore CreateListRecipeDetailStore
	}
	type args struct {
		ctx  context.Context
		data *productmodel.FoodCreate
	}

	mockFood := new(mockCreateFoodStore)
	mockCategoryFood := new(mockCreateCategoryFoodStore)
	mockCategory := new(mockUpdateCategoryStore)
	mockSizeFood := new(mockCreateSizeFoodStore)
	mockRecipeStore := new(mockCreateRecipeStore)
	mockRecipeDetailStore := new(mockCreateListRecipeDetailStore)

	fakeFoodCreate := &productmodel.FoodCreate{
		ProductCreate: &productmodel.ProductCreate{
			Name:         "Food1",
			Description:  "Description1",
			CookingGuide: "",
		},
		Categories: []string{
			"Category001",
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
			name: "CreateFood failed due to recipe store error",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodCreate,
			},
			mock: func() {
				mockFood.
					On("CreateFood",
						context.Background(),
						fakeFoodCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "CreateFood successfully",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodCreate,
			},
			mock: func() {
				mockFood.
					On("CreateFood",
						context.Background(),
						fakeFoodCreate).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := repo.CreateFood(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createFoodRepo_HandleCategoryFood(t *testing.T) {
	type fields struct {
		foodStore         CreateFoodStore
		categoryFoodStore CreateCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     CreateSizeFoodStore
		recipeStore       CreateRecipeStore
		recipeDetailStore CreateListRecipeDetailStore
	}
	type args struct {
		ctx    context.Context
		foodId string
		data   *productmodel.FoodCreate
	}

	mockFood := new(mockCreateFoodStore)
	mockCategoryFood := new(mockCreateCategoryFoodStore)
	mockCategory := new(mockUpdateCategoryStore)
	mockSizeFood := new(mockCreateSizeFoodStore)
	mockRecipeStore := new(mockCreateRecipeStore)
	mockRecipeDetailStore := new(mockCreateListRecipeDetailStore)

	fakeFoodCreate := &productmodel.FoodCreate{
		ProductCreate: &productmodel.ProductCreate{
			Name:         "Food1",
			Description:  "Description1",
			CookingGuide: "",
		},
		Categories: []string{
			"Category001",
		},
	}
	categoryFoodCreate := categoryfoodmodel.CategoryFoodCreate{
		FoodId:     "Food1",
		CategoryId: "Category001",
	}
	amount := 1
	categoryUpdateAmountProduct := categorymodel.CategoryUpdateAmountProduct{
		AmountProduct: &amount,
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
			name: "Create Food failed because can not create category food",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "Food1",
				data:   fakeFoodCreate,
			},
			mock: func() {
				mockCategoryFood.
					On("CreateCategoryFood",
						context.Background(),
						&categoryFoodCreate).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create Food failed because can not update amount product for category",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "Food1",
				data:   fakeFoodCreate,
			},
			mock: func() {
				mockCategoryFood.
					On("CreateCategoryFood",
						context.Background(),
						&categoryFoodCreate).
					Return(nil).
					Once()

				mockCategory.
					On("UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&categoryUpdateAmountProduct).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Create Food successfully",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:    context.Background(),
				foodId: "Food1",
				data:   fakeFoodCreate,
			},
			mock: func() {
				mockCategoryFood.
					On("CreateCategoryFood",
						context.Background(),
						&categoryFoodCreate).
					Return(nil).
					Once()

				mockCategory.
					On("UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&categoryUpdateAmountProduct).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := repo.HandleCategoryFood(tt.args.ctx, tt.args.foodId, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleCategoryFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleCategoryFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_createFoodRepo_HandleSizeFood(t *testing.T) {
	type fields struct {
		foodStore         CreateFoodStore
		categoryFoodStore CreateCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     CreateSizeFoodStore
		recipeStore       CreateRecipeStore
		recipeDetailStore CreateListRecipeDetailStore
	}
	type args struct {
		ctx  context.Context
		data *productmodel.FoodCreate
	}

	mockFood := new(mockCreateFoodStore)
	mockCategoryFood := new(mockCreateCategoryFoodStore)
	mockCategory := new(mockUpdateCategoryStore)
	mockSizeFood := new(mockCreateSizeFoodStore)
	mockRecipeStore := new(mockCreateRecipeStore)
	mockRecipeDetailStore := new(mockCreateListRecipeDetailStore)

	fakeFoodCreate := &productmodel.FoodCreate{
		ProductCreate: &productmodel.ProductCreate{
			Name:         "Food1",
			Description:  "Description1",
			CookingGuide: "",
		},
		Categories: []string{
			"Category001",
		},
		Sizes: []sizefoodmodel.SizeFoodCreate{
			{
				FoodId: "Food001",
				SizeId: "Size001",
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
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
			name: "Handle size food failed because can not create recipe",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodCreate,
			},
			mock: func() {
				mockRecipeStore.
					On(
						"CreateRecipe",
						context.Background(),
						fakeFoodCreate.Sizes[0].Recipe).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size food failed because can not create size food",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodCreate,
			},
			mock: func() {
				mockRecipeStore.
					On(
						"CreateRecipe",
						context.Background(),
						fakeFoodCreate.Sizes[0].Recipe).
					Return(nil).
					Once()

				mockSizeFood.
					On(
						"CreateSizeFood",
						context.Background(),
						&fakeFoodCreate.Sizes[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size food failed because can not create recipe details",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodCreate,
			},
			mock: func() {
				mockRecipeStore.
					On(
						"CreateRecipe",
						context.Background(),
						fakeFoodCreate.Sizes[0].Recipe).
					Return(nil).
					Once()

				mockSizeFood.
					On(
						"CreateSizeFood",
						context.Background(),
						&fakeFoodCreate.Sizes[0]).
					Return(nil).
					Once()

				mockRecipeDetailStore.
					On(
						"CreateListRecipeDetail",
						context.Background(),
						fakeFoodCreate.Sizes[0].Recipe.Details).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size food successfully",
			fields: fields{
				foodStore:         mockFood,
				categoryFoodStore: mockCategoryFood,
				categoryStore:     mockCategory,
				sizeFoodStore:     mockSizeFood,
				recipeStore:       mockRecipeStore,
				recipeDetailStore: mockRecipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				data: fakeFoodCreate,
			},
			mock: func() {
				mockRecipeStore.
					On(
						"CreateRecipe",
						context.Background(),
						fakeFoodCreate.Sizes[0].Recipe).
					Return(nil).
					Once()

				mockSizeFood.
					On(
						"CreateSizeFood",
						context.Background(),
						&fakeFoodCreate.Sizes[0]).
					Return(nil).
					Once()

				mockRecipeDetailStore.
					On(
						"CreateListRecipeDetail",
						context.Background(),
						fakeFoodCreate.Sizes[0].Recipe.Details).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &createFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}
			tt.mock()

			err := repo.HandleSizeFood(tt.args.ctx, tt.args.data)
			if tt.wantErr {
				assert.NotNil(t, err, "HandleSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
