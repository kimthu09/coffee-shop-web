package productrepo

import (
	"coffee_shop_management_backend/module/product/productmodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockUpdateToppingStore struct {
	mock.Mock
}

func (m *mockUpdateToppingStore) FindTopping(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*productmodel.Topping, error) {
	args := m.Called(ctx, conditions, moreKeys)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*productmodel.Topping), args.Error(1)
}

func (m *mockUpdateToppingStore) UpdateTopping(
	ctx context.Context,
	id string,
	data *productmodel.ToppingUpdateInfo) error {
	args := m.Called(ctx, id, data)
	return args.Error(0)
}

func TestNewUpdateToppingRepo(t *testing.T) {
	type args struct {
		toppingStore      UpdateToppingStore
		recipeDetailStore UpdateRecipeDetailStore
	}

	toppingStore := new(mockUpdateToppingStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	tests := []struct {
		name string
		args args
		want *updateToppingRepo
	}{
		{
			name: "Create object has type updateFoodRepo",
			args: args{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			want: &updateToppingRepo{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateToppingRepo(
				tt.args.toppingStore,
				tt.args.recipeDetailStore,
			)
			assert.Equal(t, tt.want, got, "NewUpdateToppingRepo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_updateToppingRepo_FindRecipeDetails(t *testing.T) {
	type fields struct {
		toppingStore      UpdateToppingStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx      context.Context
		recipeId string
	}

	toppingStore := new(mockUpdateToppingStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	recipeId := "Recipe001"
	recipeDetails := []recipedetailmodel.RecipeDetail{
		{
			RecipeId: "Recipe001",
		},
	}
	var moreKeys []string
	mockErr := errors.New(mock.Anything)

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
				toppingStore:      toppingStore,
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
				toppingStore:      toppingStore,
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
			repo := &updateToppingRepo{
				toppingStore:      tt.fields.toppingStore,
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

func Test_updateToppingRepo_FindTopping(t *testing.T) {
	type fields struct {
		toppingStore      UpdateToppingStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx context.Context
		id  string
	}

	toppingStore := new(mockUpdateToppingStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	toppingId := "Topping001"
	var moreKeys []string
	topping := productmodel.Topping{}

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *productmodel.Topping
		wantErr bool
	}{
		{
			name: "Find topping failed",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx: context.Background(),
				id:  toppingId,
			},
			mock: func() {
				toppingStore.
					On(
						"FindTopping",
						context.Background(),
						map[string]interface{}{"id": toppingId},
						moreKeys).
					Return(nil, mockErr).
					Once()
			},
			want:    &topping,
			wantErr: true,
		},
		{
			name: "Find topping successfully",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx: context.Background(),
				id:  toppingId,
			},
			mock: func() {
				toppingStore.
					On(
						"FindTopping",
						context.Background(),
						map[string]interface{}{"id": toppingId},
						moreKeys).
					Return(&topping, nil).
					Once()
			},
			want:    &topping,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateToppingRepo{
				toppingStore:      tt.fields.toppingStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			got, err := repo.FindTopping(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.NotNil(t, err, "FindTopping() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindFood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateToppingRepo_UpdateRecipeDetailsOfRecipe(t *testing.T) {
	type fields struct {
		toppingStore      UpdateToppingStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx                  context.Context
		recipeId             string
		deletedRecipeDetails []recipedetailmodel.RecipeDetail
		updatedRecipeDetails []recipedetailmodel.RecipeDetailUpdate
		createdRecipeDetails []recipedetailmodel.RecipeDetailCreate
	}

	toppingStore := new(mockUpdateToppingStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	recipeId := "Recipe002"
	deletedRecipeDetails := []recipedetailmodel.RecipeDetail{
		{
			RecipeId:     "Recipe002",
			IngredientId: "Ing001",
		},
	}
	updatedRecipeDetails := []recipedetailmodel.RecipeDetailUpdate{
		{
			IngredientId: "Ing002",
			AmountNeed:   20,
		},
	}
	createdRecipeDetails := []recipedetailmodel.RecipeDetailCreate{
		{
			RecipeId:     "Recipe002",
			IngredientId: "Ing003",
			AmountNeed:   30,
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
			name: "Handle recipe detail failed because can not delete recipe detail",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                  context.Background(),
				recipeId:             recipeId,
				deletedRecipeDetails: deletedRecipeDetails,
				updatedRecipeDetails: updatedRecipeDetails,
				createdRecipeDetails: createdRecipeDetails,
			},
			mock: func() {
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
			name: "Handle recipe detail failed because can not update recipe detail",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                  context.Background(),
				recipeId:             recipeId,
				deletedRecipeDetails: deletedRecipeDetails,
				updatedRecipeDetails: updatedRecipeDetails,
				createdRecipeDetails: createdRecipeDetails,
			},
			mock: func() {
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
						&updatedRecipeDetails[0]).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle recipe detail failed because can not create recipe detail",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                  context.Background(),
				recipeId:             recipeId,
				deletedRecipeDetails: deletedRecipeDetails,
				updatedRecipeDetails: updatedRecipeDetails,
				createdRecipeDetails: createdRecipeDetails,
			},
			mock: func() {
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
						&updatedRecipeDetails[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						createdRecipeDetails).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Handle recipe detail successfully",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:                  context.Background(),
				recipeId:             recipeId,
				deletedRecipeDetails: deletedRecipeDetails,
				updatedRecipeDetails: updatedRecipeDetails,
				createdRecipeDetails: createdRecipeDetails,
			},
			mock: func() {
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
						&updatedRecipeDetails[0]).
					Return(nil).
					Once()

				recipeDetailStore.
					On("CreateListRecipeDetail",
						context.Background(),
						createdRecipeDetails).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateToppingRepo{
				toppingStore:      tt.fields.toppingStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := repo.UpdateRecipeDetailsOfRecipe(
				tt.args.ctx,
				tt.args.recipeId,
				tt.args.deletedRecipeDetails,
				tt.args.updatedRecipeDetails,
				tt.args.createdRecipeDetails,
			)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateRecipeDetailsOfRecipe() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateRecipeDetailsOfRecipe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateToppingRepo_UpdateTopping(t *testing.T) {
	type fields struct {
		toppingStore      UpdateToppingStore
		recipeDetailStore UpdateRecipeDetailStore
	}
	type args struct {
		ctx  context.Context
		id   string
		data *productmodel.ToppingUpdateInfo
	}

	toppingStore := new(mockUpdateToppingStore)
	recipeDetailStore := new(mockUpdateRecipeDetailStore)

	updatedName := "Updated Name"
	updatedDescription := "Updated Description"
	updatedCookingGuide := "Updated Description"
	updatedPrice := 10
	updatedCost := 5
	updateData := &productmodel.ToppingUpdateInfo{
		ProductUpdateInfo: &productmodel.ProductUpdateInfo{
			Name:         &updatedName,
			Description:  &updatedDescription,
			CookingGuide: &updatedCookingGuide,
		},
		Cost:  &updatedCost,
		Price: &updatedPrice,
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
			name: "Update topping failed",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				id:   "Topping001",
				data: updateData,
			},
			mock: func() {
				toppingStore.
					On(
						"UpdateTopping",
						context.Background(),
						"Topping001",
						updateData).
					Return(mockErr).
					Once()
			},
			wantErr: true,
		},
		{
			name: "Update topping successfully",
			fields: fields{
				toppingStore:      toppingStore,
				recipeDetailStore: recipeDetailStore,
			},
			args: args{
				ctx:  context.Background(),
				id:   "Topping001",
				data: updateData,
			},
			mock: func() {
				toppingStore.
					On(
						"UpdateTopping",
						context.Background(),
						"Topping001",
						updateData).
					Return(nil).
					Once()
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &updateToppingRepo{
				toppingStore:      tt.fields.toppingStore,
				recipeDetailStore: tt.fields.recipeDetailStore,
			}

			tt.mock()

			err := repo.UpdateTopping(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
