package recipemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecipeCreate_TableName(t *testing.T) {
	type fields struct {
		Id      string
		Details []recipedetailmodel.RecipeDetailCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of RecipeCreate successfully",
			fields: fields{
				Id:      "",
				Details: nil,
			},
			want: common.TableRecipe,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recipe := &RecipeCreate{
				Id:      tt.fields.Id,
				Details: tt.fields.Details,
			}

			got := recipe.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestRecipeCreate_Validate(t *testing.T) {
	type fields struct {
		Id      string
		Details []recipedetailmodel.RecipeDetailCreate
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "RecipeCreate is invalid with Details slice of length 0",
			fields: fields{
				Details: make([]recipedetailmodel.RecipeDetailCreate, 0),
			},
			wantErr: true,
		},
		{
			name: "RecipeCreate is invalid with empty Details",
			fields: fields{
				Details: nil,
			},
			wantErr: true,
		},
		{
			name: "RecipeCreate is invalid with invalid detail",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						IngredientId: "",
						AmountNeed:   10.0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "RecipeCreate is invalid with duplicate IngredientId",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						IngredientId: "ingredient1",
						AmountNeed:   10.0,
					},
					{
						IngredientId: "ingredient1",
						AmountNeed:   15.0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "RecipeCreate is valid",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailCreate{
					{
						IngredientId: "ingredient1",
						AmountNeed:   10.0,
					},
					{
						IngredientId: "ingredient2",
						AmountNeed:   15.0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &RecipeCreate{
				Id:      tt.fields.Id,
				Details: tt.fields.Details,
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
