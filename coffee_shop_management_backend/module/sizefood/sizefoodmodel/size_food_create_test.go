package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeFoodCreate_TableName(t *testing.T) {
	type fields struct {
		FoodId   string
		SizeId   string
		Name     string
		Cost     int
		Price    int
		RecipeId string
		Recipe   *recipemodel.RecipeCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of SizeFoodCreate successfully",
			fields: fields{
				FoodId:   "",
				SizeId:   "",
				Name:     "",
				Cost:     0,
				Price:    0,
				RecipeId: "",
				Recipe:   nil,
			},
			want: common.TableSizeFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sizeFood := &SizeFoodCreate{
				FoodId:   tt.fields.FoodId,
				SizeId:   tt.fields.SizeId,
				Name:     tt.fields.Name,
				Cost:     tt.fields.Cost,
				Price:    tt.fields.Price,
				RecipeId: tt.fields.RecipeId,
				Recipe:   tt.fields.Recipe,
			}

			got := sizeFood.TableName()

			assert.Equal(
				t,
				tt.want,
				got,
				"TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestSizeFoodCreate_Validate(t *testing.T) {
	type fields struct {
		FoodId   string
		SizeId   string
		Name     string
		Cost     int
		Price    int
		RecipeId string
		Recipe   *recipemodel.RecipeCreate
	}

	invalidRecipe := recipemodel.RecipeCreate{
		Details: nil,
	}
	validRecipe := recipemodel.RecipeCreate{
		Details: []recipedetailmodel.RecipeDetailCreate{
			{
				IngredientId: "0123",
				AmountNeed:   10,
			},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "SizeFoodCreate is invalid with empty name",
			fields: fields{
				Name:   "",
				Cost:   10,
				Price:  15,
				Recipe: &validRecipe,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodCreate is invalid with negative cost",
			fields: fields{
				Name:   "Regular",
				Cost:   -10,
				Price:  15,
				Recipe: &validRecipe,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodCreate is invalid with negative price",
			fields: fields{
				Name:   "Regular",
				Cost:   10,
				Price:  -15,
				Recipe: &validRecipe,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodCreate is invalid with empty recipe",
			fields: fields{
				Name:   "Regular",
				Cost:   10,
				Price:  15,
				Recipe: nil,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodCreate is invalid with invalid recipe",
			fields: fields{
				Name:   "Regular",
				Cost:   10,
				Price:  15,
				Recipe: &invalidRecipe,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodCreate is valid",
			fields: fields{
				Name:   "Regular",
				Cost:   10,
				Price:  15,
				Recipe: &validRecipe,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &SizeFoodCreate{
				FoodId:   tt.fields.FoodId,
				SizeId:   tt.fields.SizeId,
				Name:     tt.fields.Name,
				Cost:     tt.fields.Cost,
				Price:    tt.fields.Price,
				RecipeId: tt.fields.RecipeId,
				Recipe:   tt.fields.Recipe,
			}
			err := data.Validate()

			if tt.wantErr {
				assert.NotNil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
