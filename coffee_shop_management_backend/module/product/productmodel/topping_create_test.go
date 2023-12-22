package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToppingCreate_TableName(t *testing.T) {
	type fields struct {
		ProductCreate *ProductCreate
		Cost          int
		Price         int
		RecipeId      string
		Recipe        *recipemodel.RecipeCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ToppingCreate successfully",
			fields: fields{
				ProductCreate: &ProductCreate{},
				Cost:          10.0,
				Price:         20.0,
				RecipeId:      "Recipe123",
				Recipe:        &recipemodel.RecipeCreate{},
			},
			want: common.TableTopping,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topping := &ToppingCreate{
				ProductCreate: tt.fields.ProductCreate,
				Cost:          tt.fields.Cost,
				Price:         tt.fields.Price,
				RecipeId:      tt.fields.RecipeId,
				Recipe:        tt.fields.Recipe,
			}
			got := topping.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestToppingCreate_Validate(t *testing.T) {
	type fields struct {
		ProductCreate *ProductCreate
		Cost          int
		Price         int
		RecipeId      string
		Recipe        *recipemodel.RecipeCreate
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ToppingCreate is valid with valid data",
			fields: fields{
				ProductCreate: &ProductCreate{
					Id:           nil,
					Name:         "Topping001",
					Description:  "",
					CookingGuide: "",
				},
				Cost:  10,
				Price: 20,
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ToppingCreate is invalid with empty product information",
			fields: fields{
				ProductCreate: nil,
				Cost:          10,
				Price:         20,
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ToppingCreate is invalid with invalid product information",
			fields: fields{
				ProductCreate: &ProductCreate{
					Id:           nil,
					Name:         "",
					Description:  "",
					CookingGuide: "",
				},
				Cost:  10,
				Price: 20,
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ToppingCreate is invalid with negative cost",
			fields: fields{
				ProductCreate: &ProductCreate{
					Id:           nil,
					Name:         "Topping001",
					Description:  "",
					CookingGuide: "",
				},
				Cost:  -10,
				Price: 20,
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ToppingCreate is invalid with negative price",
			fields: fields{
				ProductCreate: &ProductCreate{
					Id:           nil,
					Name:         "Topping001",
					Description:  "",
					CookingGuide: "",
				},
				Cost:  10,
				Price: -20,
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ToppingCreate is invalid with recipe not exist",
			fields: fields{
				ProductCreate: &ProductCreate{
					Id:           nil,
					Name:         "Topping001",
					Description:  "",
					CookingGuide: "",
				},
				Cost:   10,
				Price:  20,
				Recipe: nil,
			},
			wantErr: true,
		},
		{
			name: "ToppingCreate is invalid with invalid recipe",
			fields: fields{
				ProductCreate: &ProductCreate{
					Id:           nil,
					Name:         "Topping001",
					Description:  "",
					CookingGuide: "",
				},
				Cost:  10,
				Price: 20,
				Recipe: &recipemodel.RecipeCreate{
					Details: []recipedetailmodel.RecipeDetailCreate{
						{
							IngredientId: "",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ToppingCreate{
				ProductCreate: tt.fields.ProductCreate,
				Cost:          tt.fields.Cost,
				Price:         tt.fields.Price,
				RecipeId:      tt.fields.RecipeId,
				Recipe:        tt.fields.Recipe,
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
