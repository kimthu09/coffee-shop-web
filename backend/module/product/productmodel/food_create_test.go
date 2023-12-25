package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoodCreate_TableName(t *testing.T) {
	type fields struct {
		ProductCreate *ProductCreate
		Categories    []string
		Sizes         []sizefoodmodel.SizeFoodCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of FoodCreate successfully",
			fields: fields{
				ProductCreate: nil,
				Categories:    nil,
				Sizes:         nil,
			},
			want: common.TableFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			food := &FoodCreate{
				ProductCreate: tt.fields.ProductCreate,
				Categories:    tt.fields.Categories,
				Sizes:         tt.fields.Sizes,
			}
			got := food.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestFoodCreate_Validate(t *testing.T) {
	type fields struct {
		ProductCreate *ProductCreate
		Categories    []string
		Sizes         []sizefoodmodel.SizeFoodCreate
	}
	emptyCategories := make([]string, 0)
	duplicateCategories := []string{"Category1", "Category1"}
	emptySizes := make([]sizefoodmodel.SizeFoodCreate, 0)
	invalidSize := []sizefoodmodel.SizeFoodCreate{
		{
			Name: "",
		},
	}
	validSize := []sizefoodmodel.SizeFoodCreate{
		{
			Name: "Size1",
			Recipe: &recipemodel.RecipeCreate{
				Id: "Recipe1",
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						IngredientId: "Ing001",
						AmountNeed:   10,
					},
				},
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "FoodCreate is invalid with empty product information",
			fields: fields{
				ProductCreate: nil,
				Categories:    []string{"Category1"},
				Sizes:         validSize,
			},
			wantErr: true,
		},
		{
			name: "FoodCreate is invalid with invalid product information",
			fields: fields{
				ProductCreate: &ProductCreate{},
				Categories:    []string{"Category1"},
				Sizes:         validSize,
			},
			wantErr: true,
		},
		{
			name: "FoodCreate is invalid with empty categories",
			fields: fields{
				ProductCreate: &ProductCreate{
					Name: "Food Name",
				},
				Categories: emptyCategories,
				Sizes:      validSize,
			},
			wantErr: true,
		},
		{
			name: "FoodCreate is invalid with duplicate categories",
			fields: fields{
				ProductCreate: &ProductCreate{
					Name: "Food Name",
				},
				Categories: duplicateCategories,
				Sizes:      validSize,
			},
			wantErr: true,
		},
		{
			name: "FoodCreate is invalid with empty sizes",
			fields: fields{
				ProductCreate: &ProductCreate{
					Name: "Food Name",
				},
				Categories: []string{"Category1"},
				Sizes:      emptySizes,
			},
			wantErr: true,
		},
		{
			name: "FoodCreate is invalid with invalid size",
			fields: fields{
				ProductCreate: &ProductCreate{
					Name: "Food Name",
				},
				Categories: []string{"Category1"},
				Sizes:      invalidSize,
			},
			wantErr: true,
		},
		{
			name: "FoodCreate is valid",
			fields: fields{
				ProductCreate: &ProductCreate{
					Name: "Food Name",
				},
				Categories: []string{"Category1"},
				Sizes:      validSize,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			food := &FoodCreate{
				ProductCreate: tt.fields.ProductCreate,
				Categories:    tt.fields.Categories,
				Sizes:         tt.fields.Sizes,
			}
			err := food.Validate()

			if tt.wantErr {
				assert.NotNil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
