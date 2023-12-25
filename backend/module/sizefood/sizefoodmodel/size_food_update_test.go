package sizefoodmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeFoodUpdate_TableName(t *testing.T) {
	type fields struct {
		SizeId   *string
		Name     *string
		Cost     *int
		Price    *int
		RecipeId *string
		Recipe   *recipemodel.RecipeUpdate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of SizeFoodUpdate successfully",
			fields: fields{
				SizeId:   nil,
				Name:     nil,
				Cost:     nil,
				Price:    nil,
				RecipeId: nil,
				Recipe:   nil,
			},
			want: common.TableSizeFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sizeFood := &SizeFoodUpdate{
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

func TestSizeFoodUpdate_Validate(t *testing.T) {
	type fields struct {
		SizeId   *string
		Name     *string
		Cost     *int
		Price    *int
		RecipeId *string
		Recipe   *recipemodel.RecipeUpdate
	}

	emptyName := ""
	name := "Name"

	negativeNumber := -15
	number := 15

	invalidRecipeUpdate := recipemodel.RecipeUpdate{
		Details: nil,
	}
	validRecipeUpdate := recipemodel.RecipeUpdate{
		Details: []recipedetailmodel.RecipeDetailUpdate{
			{
				IngredientId: "1234567",
				AmountNeed:   10.0,
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "SizeFoodUpdate is invalid with empty name",
			fields: fields{
				Name: &emptyName,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodUpdate is invalid with negative cost",
			fields: fields{
				Name: nil,
				Cost: &negativeNumber,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodUpdate is invalid with negative price",
			fields: fields{
				Name:  nil,
				Cost:  nil,
				Price: &negativeNumber,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodUpdate is invalid with invalid recipe",
			fields: fields{
				Name:   nil,
				Cost:   nil,
				Price:  nil,
				Recipe: &invalidRecipeUpdate,
			},
			wantErr: true,
		},
		{
			name: "SizeFoodUpdate is valid",
			fields: fields{
				Name:   &name,
				Cost:   &number,
				Price:  &number,
				Recipe: &validRecipeUpdate,
			},
			wantErr: false,
		},
		{
			name: "SizeFoodUpdate is valid with nil fields",
			fields: fields{
				Name:   nil,
				Cost:   nil,
				Price:  nil,
				Recipe: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &SizeFoodUpdate{
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
