package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoodUpdateInfo_TableName(t *testing.T) {
	type fields struct {
		ProductUpdate *ProductUpdateInfo
		Categories    *[]string
		Sizes         *[]sizefoodmodel.SizeFoodUpdate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ProductUpdateInfo successfully",
			fields: fields{
				ProductUpdate: nil,
				Categories:    nil,
				Sizes:         nil,
			},
			want: common.TableFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			food := &FoodUpdateInfo{
				ProductUpdateInfo: tt.fields.ProductUpdate,
				Categories:        tt.fields.Categories,
				Sizes:             tt.fields.Sizes,
			}
			got := food.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestFoodUpdateInfo_Validate(t *testing.T) {
	type fields struct {
		ProductUpdateInfo *ProductUpdateInfo
		Categories        *[]string
		Sizes             *[]sizefoodmodel.SizeFoodUpdate
	}
	emptyName := ""
	emptyCategories := &[]string{}
	duplicatedCategories := &[]string{"Category1", "Category1"}
	emptySizes := &[]sizefoodmodel.SizeFoodUpdate{}
	invalidSize := &[]sizefoodmodel.SizeFoodUpdate{
		{
			Name: &emptyName,
		},
	}
	sizeName := "Size name"
	sizeId := "Size001"
	duplicatedSizes := &[]sizefoodmodel.SizeFoodUpdate{
		{
			SizeId: &sizeId,
			Recipe: &recipemodel.RecipeUpdate{
				Details: []recipedetailmodel.RecipeDetailUpdate{
					{
						IngredientId: "Ing001",
						AmountNeed:   10,
					},
				},
			},
		},
		{
			SizeId: &sizeId,
			Recipe: &recipemodel.RecipeUpdate{
				Details: []recipedetailmodel.RecipeDetailUpdate{
					{
						IngredientId: "Ing001",
						AmountNeed:   10,
					},
				},
			},
		},
	}
	validSize := &[]sizefoodmodel.SizeFoodUpdate{
		{
			Name: &sizeName,
			Recipe: &recipemodel.RecipeUpdate{
				Details: []recipedetailmodel.RecipeDetailUpdate{
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
			name: "ProductUpdateInfo is invalid with invalid product info",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name: &emptyName,
				},
				Categories: &[]string{"Category1"},
				Sizes:      validSize,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateInfo is invalid with empty categories",
			fields: fields{
				ProductUpdateInfo: nil,
				Categories:        emptyCategories,
				Sizes:             validSize,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateInfo is invalid with duplicate categories",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{},
				Categories:        duplicatedCategories,
				Sizes:             validSize,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateInfo is invalid with empty sizes",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{},
				Categories:        &[]string{"Category1"},
				Sizes:             emptySizes,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateInfo is invalid with invalid size",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{},
				Categories:        &[]string{"Category1"},
				Sizes:             invalidSize,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateInfo is invalid with duplicate sizes",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{},
				Categories:        &[]string{"Category1"},
				Sizes:             duplicatedSizes,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateInfo is valid",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{},
				Categories:        &[]string{"Category1"},
				Sizes:             validSize,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foodUpdate := &FoodUpdateInfo{
				ProductUpdateInfo: tt.fields.ProductUpdateInfo,
				Categories:        tt.fields.Categories,
				Sizes:             tt.fields.Sizes,
			}
			err := foodUpdate.Validate()

			if tt.wantErr {
				assert.NotNil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Validate() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
