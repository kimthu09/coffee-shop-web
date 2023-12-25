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

type mockUpdateFoodStore struct {
	mock.Mock
}

func (m *mockUpdateFoodStore) UpdateFood(
	ctx context.Context, id string,
	data *productmodel.FoodUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func (m *mockUpdateFoodStore) FindFood(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*productmodel.Food, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Food), args.Error(1)
}

type mockCreateOrDeleteCategoryFoodStore struct {
	mock.Mock
}

func (m *mockCreateOrDeleteCategoryFoodStore) FindListCategories(
	ctx context.Context,
	foodId string) ([]categorymodel.SimpleCategoryWithId, error) {
	args := m.Called(ctx, foodId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]categorymodel.SimpleCategoryWithId), args.Error(1)
}

func (m *mockCreateOrDeleteCategoryFoodStore) CreateCategoryFood(ctx context.Context, data *categoryfoodmodel.CategoryFoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockCreateOrDeleteCategoryFoodStore) DeleteCategoryFood(ctx context.Context, conditions map[string]interface{}) error {
	args := m.Called(ctx, conditions)
	return args.Error(0)
}

type mockUpdateSizeFoodStore struct {
	mock.Mock
}

func (m *mockUpdateSizeFoodStore) FindListSizeFood(
	ctx context.Context,
	foodId string) ([]sizefoodmodel.SizeFood, error) {
	args := m.Called(ctx, foodId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]sizefoodmodel.SizeFood), args.Error(1)
}

func (m *mockUpdateSizeFoodStore) CreateSizeFood(
	ctx context.Context,
	data *sizefoodmodel.SizeFoodCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockUpdateSizeFoodStore) DeleteSizeFood(
	ctx context.Context,
	conditions map[string]interface{}) error {
	args := m.Called(ctx, conditions)
	return args.Error(0)
}

func (m *mockUpdateSizeFoodStore) UpdateSizeFood(
	ctx context.Context,
	foodId string,
	sizeId string,
	data *sizefoodmodel.SizeFoodUpdate) error {
	args := m.Called(ctx, foodId, sizeId, data)
	return args.Error(0)
}

type mockUpdateRecipeStore struct {
	mock.Mock
}

func (m *mockUpdateRecipeStore) CreateRecipe(
	ctx context.Context,
	data *recipemodel.RecipeCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockUpdateRecipeStore) DeleteRecipe(
	ctx context.Context,
	conditions map[string]interface{}) error {
	args := m.Called(ctx, conditions)
	return args.Error(0)
}

type mockUpdateRecipeDetailStore struct {
	mock.Mock
}

func (m *mockUpdateRecipeDetailStore) UpdateRecipeDetail(ctx context.Context, idRecipe, idIngredient string, data *recipedetailmodel.RecipeDetailUpdate) error {
	args := m.Called(ctx, idRecipe, idIngredient, data)
	return args.Error(0)
}

func (m *mockUpdateRecipeDetailStore) CreateListRecipeDetail(ctx context.Context, data []recipedetailmodel.RecipeDetailCreate) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

func (m *mockUpdateRecipeDetailStore) FindListRecipeDetail(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) ([]recipedetailmodel.RecipeDetail, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]recipedetailmodel.RecipeDetail), args.Error(1)
}

func (m *mockUpdateRecipeDetailStore) DeleteRecipeDetail(ctx context.Context, conditions map[string]interface{}) error {
	args := m.Called(ctx, conditions)
	return args.Error(0)
}

func TestNewUpdateFoodRepo(t *testing.T) {
	type args struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	tests := []struct {
		name string
		args args
		want *updateFoodRepo
	}{
		{
			name: "Create object has type updateFoodRepo",
			args: args{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			want: &updateFoodRepo{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateFoodRepo(
				tt.args.foodStore,
				tt.args.categoryFoodStore,
				tt.args.categoryStore,
				tt.args.sizeFoodStore,
				tt.args.recipeStore,
				tt.args.recipeDetailStore,
			)
			assert.Equal(t, tt.want, got, "NewUpdateFoodRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateFoodBiz_FindCategories(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx    context.Context
		foodId string
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	ctx := context.Background()
	foodId := "Food001"

	categories := []categorymodel.SimpleCategoryWithId{
		{CategoryId: "category1"},
	}
	mockErr := assert.AnError

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []categorymodel.SimpleCategoryWithId
		wantErr bool
	}{
		{
			name: "FindListCategories failed",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:    ctx,
				foodId: foodId,
			},
			mock: func() {
				categoryFoodStore.
					On("FindListCategories", ctx, foodId).
					Return(nil, mockErr).
					Once()
			},
			want:    categories,
			wantErr: true,
		},
		{
			name: "FindListCategories successful",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:    ctx,
				foodId: foodId,
			},
			mock: func() {
				categoryFoodStore.
					On("FindListCategories", ctx, foodId).
					Return(categories, nil).
					Once()
			},
			want:    categories,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			got, err := biz.FindCategories(tt.args.ctx, tt.args.foodId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindCategories() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindCategories() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateFoodBiz_FindFood(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx context.Context
		id  string
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	ctx := context.Background()
	foodId := "Food001"
	var moreKeys []string
	mockErr := errors.New(mock.Anything)
	food := productmodel.Food{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Food
		wantErr bool
	}{
		{
			name: "Find food failed",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx: ctx,
				id:  foodId,
			},
			mock: func() {
				store.
					On(
						"FindFood",
						ctx,
						map[string]interface{}{"id": "Food001"},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    &food,
			wantErr: true,
		},
		{
			name: "FindFood successful",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx: ctx,
				id:  foodId,
			},
			mock: func() {
				store.
					On(
						"FindFood",
						ctx,
						map[string]interface{}{"id": "Food001"},
						moreKeys).
					Return(&food, nil).
					Once()
			},
			want:    &food,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			got, err := biz.FindFood(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.NotNil(t, err, "FindFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindFood() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateFoodBiz_FindRecipeDetails(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	recipeId := "Recipe001"
	recipeDetails := []recipedetailmodel.RecipeDetail{
		{
			RecipeId: "Recipe001",
		},
	}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)

	type args struct {
		ctx      context.Context
		recipeId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []recipedetailmodel.RecipeDetail
		wantErr bool
	}{
		{
			name: "Find list recipe details failed",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:      context.Background(),
				recipeId: recipeId,
			},
			mock: func() {
				recipeDetailStore.
					On(
						"FindListRecipeDetail",
						context.Background(),
						map[string]interface{}{"recipeId": "Recipe001"},
						moreKeys,
					).
					Return(nil, mockErr).
					Once()
			},
			want:    recipeDetails,
			wantErr: true,
		},
		{
			name: "Find list recipe details successfully",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:      context.Background(),
				recipeId: recipeId,
			},
			mock: func() {
				recipeDetailStore.
					On(
						"FindListRecipeDetail",
						context.Background(),
						map[string]interface{}{"recipeId": "Recipe001"},
						moreKeys,
					).
					Return(recipeDetails, nil).
					Once()
			},
			want:    recipeDetails,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			got, err := repo.FindRecipeDetails(tt.args.ctx, tt.args.recipeId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindRecipeDetails() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindRecipeDetails() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindRecipeDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateFoodRepo_FindSizeFoods(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx    context.Context
		foodId string
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	ctx := context.Background()
	foodId := "Food001"

	sizeFoods := []sizefoodmodel.SizeFood{
		{
			FoodId: "Food001",
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []sizefoodmodel.SizeFood
		wantErr bool
	}{
		{
			name: "Find list size food failed",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:    ctx,
				foodId: foodId,
			},
			mock: func() {
				sizeFoodStore.
					On("FindListSizeFood", ctx, foodId).
					Return(nil, mockErr).
					Once()
			},
			want:    sizeFoods,
			wantErr: true,
		},
		{
			name: "Find list size food successful",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:    ctx,
				foodId: foodId,
			},
			mock: func() {
				sizeFoodStore.
					On("FindListSizeFood", ctx, foodId).
					Return(sizeFoods, nil).
					Once()
			},
			want:    sizeFoods,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			got, err := repo.FindSizeFoods(tt.args.ctx, tt.args.foodId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindSizeFoods() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindSizeFoods() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindSizeFoods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateFoodBiz_HandleCategory(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx                 context.Context
		foodId              string
		deletedCategoryFood []categorymodel.SimpleCategoryWithId
		createdCategoryFood []categorymodel.SimpleCategoryWithId
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	deletedCategoryFood := []categorymodel.SimpleCategoryWithId{
		{
			CategoryId: "Category001",
		},
	}
	createdCategoryFood := []categorymodel.SimpleCategoryWithId{
		{
			CategoryId: "Category002",
		},
	}
	amountMinus := -1
	amountPlus := 1
	updateModel1 := categorymodel.CategoryUpdateAmountProduct{AmountProduct: &amountMinus}
	updateModel2 := categorymodel.CategoryUpdateAmountProduct{AmountProduct: &amountPlus}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Handle category failed because can not minus 1 total amount product of category",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                 context.Background(),
				foodId:              "Food001",
				deletedCategoryFood: deletedCategoryFood,
				createdCategoryFood: createdCategoryFood,
			},
			mock: func() {
				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&updateModel1).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle category failed because can not plus 1 total amount product of category",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                 context.Background(),
				foodId:              "Food001",
				deletedCategoryFood: deletedCategoryFood,
				createdCategoryFood: createdCategoryFood,
			},
			mock: func() {
				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&updateModel1).
					Return(nil).
					Once()

				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category002",
						&updateModel2).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle category failed because can not delete size food",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                 context.Background(),
				foodId:              "Food001",
				deletedCategoryFood: deletedCategoryFood,
				createdCategoryFood: createdCategoryFood,
			},
			mock: func() {
				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&updateModel1).
					Return(nil).
					Once()

				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category002",
						&updateModel2).
					Return(nil).
					Once()

				categoryFoodStore.
					On(
						"DeleteCategoryFood",
						context.Background(),
						map[string]interface{}{
							"foodId":     "Food001",
							"categoryId": "Category001",
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle category failed because can not create size food",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                 context.Background(),
				foodId:              "Food001",
				deletedCategoryFood: deletedCategoryFood,
				createdCategoryFood: createdCategoryFood,
			},
			mock: func() {
				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&updateModel1).
					Return(nil).
					Once()

				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category002",
						&updateModel2).
					Return(nil).
					Once()

				categoryFoodStore.
					On(
						"DeleteCategoryFood",
						context.Background(),
						map[string]interface{}{
							"foodId":     "Food001",
							"categoryId": "Category001",
						}).
					Return(nil).
					Once()

				categoryFoodStore.
					On(
						"CreateCategoryFood",
						context.Background(),
						&categoryfoodmodel.CategoryFoodCreate{
							FoodId:     "Food001",
							CategoryId: "Category002",
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle category successfully",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                 context.Background(),
				foodId:              "Food001",
				deletedCategoryFood: deletedCategoryFood,
				createdCategoryFood: createdCategoryFood,
			},
			mock: func() {
				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category001",
						&updateModel1).
					Return(nil).
					Once()

				categoryStore.
					On(
						"UpdateAmountProductCategory",
						context.Background(),
						"Category002",
						&updateModel2).
					Return(nil).
					Once()

				categoryFoodStore.
					On(
						"DeleteCategoryFood",
						context.Background(),
						map[string]interface{}{
							"foodId":     "Food001",
							"categoryId": "Category001",
						}).
					Return(nil).
					Once()

				categoryFoodStore.
					On(
						"CreateCategoryFood",
						context.Background(),
						&categoryfoodmodel.CategoryFoodCreate{
							FoodId:     "Food001",
							CategoryId: "Category002",
						}).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := repo.HandleCategory(
				tt.args.ctx,
				tt.args.foodId,
				tt.args.deletedCategoryFood,
				tt.args.createdCategoryFood,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleCategory() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_updateFoodBiz_HandleSizeFoods(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx                     context.Context
		foodId                  string
		deletedSizeFood         []sizefoodmodel.SizeFood
		updatedSizeFood         []sizefoodmodel.SizeFoodUpdate
		mapDeletedRecipeDetails map[string][]recipedetailmodel.RecipeDetail
		mapUpdatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailUpdate
		mapCreatedRecipeDetails map[string][]recipedetailmodel.RecipeDetailCreate
		createdSizeFood         []sizefoodmodel.SizeFoodCreate
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	ctx := context.Background()
	foodId := "Food001"

	deletedSizeFood := []sizefoodmodel.SizeFood{
		{
			FoodId:   "Food001",
			SizeId:   "Size001",
			RecipeId: "Recipe001",
		},
	}
	sizeUpdate := "Size002"
	recipeUpdate := "Recipe002"
	updatedSizeFood := []sizefoodmodel.SizeFoodUpdate{
		{
			SizeId:   &sizeUpdate,
			RecipeId: &recipeUpdate,
		},
	}
	mapDeletedRecipeDetails := map[string][]recipedetailmodel.RecipeDetail{
		"Recipe002": []recipedetailmodel.RecipeDetail{
			{
				RecipeId:     "Recipe002",
				IngredientId: "Ing001",
			},
		},
	}
	mapUpdatedRecipeDetails := map[string][]recipedetailmodel.RecipeDetailUpdate{
		"Recipe002": []recipedetailmodel.RecipeDetailUpdate{
			{
				IngredientId: "Ing002",
				AmountNeed:   20,
			},
		},
	}
	mapCreatedRecipeDetails := map[string][]recipedetailmodel.RecipeDetailCreate{
		"Recipe002": []recipedetailmodel.RecipeDetailCreate{
			{
				RecipeId:     "Recipe002",
				IngredientId: "Ing003",
				AmountNeed:   30,
			},
		},
	}
	createdSizeFood := []sizefoodmodel.SizeFoodCreate{
		{
			FoodId: "Food001",
			SizeId: "Size003",
			Recipe: &recipemodel.RecipeCreate{
				Id: "Recipe003",
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						RecipeId:     "Recipe003",
						IngredientId: "Ing001",
						AmountNeed:   10,
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
			name: "Handle size foods failed because can not delete recipe detail",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not delete size food",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not delete size food",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not update size food",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not delete recipe detail",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not update recipe detail",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(nil).
					Once()

				recipeDetailStore.
					On("UpdateRecipeDetail",
						context.Background(),
						"Recipe002",
						"Ing002",
						&mapUpdatedRecipeDetails["Recipe002"][0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not create list recipe detail",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(nil).
					Once()

				recipeDetailStore.
					On("UpdateRecipeDetail",
						context.Background(),
						"Recipe002",
						"Ing002",
						&mapUpdatedRecipeDetails["Recipe002"][0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						mapCreatedRecipeDetails["Recipe002"]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not create recipe",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(nil).
					Once()

				recipeDetailStore.
					On("UpdateRecipeDetail",
						context.Background(),
						"Recipe002",
						"Ing002",
						&mapUpdatedRecipeDetails["Recipe002"][0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						mapCreatedRecipeDetails["Recipe002"]).
					Return(nil).
					Once()

				recipeStore.
					On("CreateRecipe",
						context.Background(),
						createdSizeFood[0].Recipe).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not create list recipe detail",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(nil).
					Once()

				recipeDetailStore.
					On("UpdateRecipeDetail",
						context.Background(),
						"Recipe002",
						"Ing002",
						&mapUpdatedRecipeDetails["Recipe002"][0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						mapCreatedRecipeDetails["Recipe002"]).
					Return(nil).
					Once()

				recipeStore.
					On("CreateRecipe",
						context.Background(),
						createdSizeFood[0].Recipe).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						createdSizeFood[0].Recipe.Details).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods failed because can not create size food",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(nil).
					Once()

				recipeDetailStore.
					On("UpdateRecipeDetail",
						context.Background(),
						"Recipe002",
						"Ing002",
						&mapUpdatedRecipeDetails["Recipe002"][0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						mapCreatedRecipeDetails["Recipe002"]).
					Return(nil).
					Once()

				recipeStore.
					On("CreateRecipe",
						context.Background(),
						createdSizeFood[0].Recipe).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						createdSizeFood[0].Recipe.Details).
					Return(nil).
					Once()

				sizeFoodStore.
					On("CreateSizeFood",
						context.Background(),
						&createdSizeFood[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle size foods successfully",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                     ctx,
				foodId:                  foodId,
				deletedSizeFood:         deletedSizeFood,
				updatedSizeFood:         updatedSizeFood,
				mapDeletedRecipeDetails: mapDeletedRecipeDetails,
				mapUpdatedRecipeDetails: mapUpdatedRecipeDetails,
				mapCreatedRecipeDetails: mapCreatedRecipeDetails,
				createdSizeFood:         createdSizeFood,
			},
			mock: func() {
				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("DeleteSizeFood",
						context.Background(),
						map[string]interface{}{
							"foodId": foodId,
							"sizeId": "Size001",
						}).
					Return(nil).
					Once()

				recipeStore.
					On("DeleteRecipe",
						context.Background(),
						map[string]interface{}{
							"id": "Recipe001",
						}).
					Return(nil).
					Once()

				sizeFoodStore.
					On("UpdateSizeFood",
						context.Background(),
						foodId,
						"Size002",
						&updatedSizeFood[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("DeleteRecipeDetail",
						context.Background(),
						map[string]interface{}{
							"recipeId":     "Recipe002",
							"ingredientId": "Ing001",
						}).
					Return(nil).
					Once()

				recipeDetailStore.
					On("UpdateRecipeDetail",
						context.Background(),
						"Recipe002",
						"Ing002"+
							"",
						&mapUpdatedRecipeDetails["Recipe002"][0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						mapCreatedRecipeDetails["Recipe002"]).
					Return(nil).
					Once()

				recipeStore.
					On("CreateRecipe",
						context.Background(),
						createdSizeFood[0].Recipe).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						createdSizeFood[0].Recipe.Details).
					Return(nil).
					Once()

				sizeFoodStore.
					On("CreateSizeFood",
						context.Background(),
						&createdSizeFood[0]).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			biz := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := biz.HandleSizeFoods(
				tt.args.ctx,
				tt.args.foodId,
				tt.args.deletedSizeFood,
				tt.args.updatedSizeFood,
				tt.args.mapDeletedRecipeDetails,
				tt.args.mapUpdatedRecipeDetails,
				tt.args.mapCreatedRecipeDetails,
				tt.args.createdSizeFood,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "HandleSizeFoods() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "HandleSizeFoods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateFoodBiz_UpdateFood(t *testing.T) {
	type fields struct {
		foodStore         UpdateFoodStore
		categoryFoodStore CreateOrDeleteCategoryFoodStore
		categoryStore     UpdateCategoryStore
		sizeFoodStore     UpdateSizeFoodStore
		recipeStore       UpdateRecipeStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx  context.Context
		id   string
		data *productmodel.FoodUpdateInfo
	}

	store := new(mockUpdateFoodStore)
	categoryFoodStore := new(mockCreateOrDeleteCategoryFoodStore)
	categoryStore := new(mockUpdateCategoryStore)
	sizeFoodStore := new(mockUpdateSizeFoodStore)
	recipeStore := new(mockUpdateRecipeStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	updatedName := "Updated Name"
	updatedDescription := "Updated Description"
	updatedCookingGuide := "Updated Description"
	updateData := &productmodel.FoodUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &updatedName,
			Description:  &updatedDescription,
			CookingGuide: &updatedCookingGuide,
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
			name: "Update food failed",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				id:   "Food001",
				data: updateData,
			},
			mock: func() {
				store.
					On(
						"UpdateFood",
						context.Background(),
						"Food001",
						updateData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update food successfully",
			fields: fields{
				foodStore:         store,
				categoryFoodStore: categoryFoodStore,
				categoryStore:     categoryStore,
				sizeFoodStore:     sizeFoodStore,
				recipeStore:       recipeStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				id:   "Food001",
				data: updateData,
			},
			mock: func() {
				store.
					On(
						"UpdateFood",
						context.Background(),
						"Food001",
						updateData).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateFoodRepo{
				foodStore:         tt.fields.foodStore,
				categoryFoodStore: tt.fields.categoryFoodStore,
				categoryStore:     tt.fields.categoryStore,
				sizeFoodStore:     tt.fields.sizeFoodStore,
				recipeStore:       tt.fields.recipeStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}
			tt.mock()

			err := repo.UpdateFood(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
