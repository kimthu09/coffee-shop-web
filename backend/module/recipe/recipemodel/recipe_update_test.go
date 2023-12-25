package recipemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecipeUpdate_TableName(t *testing.T) {
	type fields struct {
		Details []recipedetailmodel.RecipeDetailUpdate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of RecipeUpdate successfully",
			fields: fields{
				Details: nil,
			},
			want: common.TableRecipe,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &RecipeUpdate{
				Details: tt.fields.Details,
			}
			assert.Equalf(t, tt.want, re.TableName(), "TableName()")
		})
	}
}

func TestRecipeUpdate_Validate(t *testing.T) {
	type fields struct {
		Details []recipedetailmodel.RecipeDetailUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "RecipeUpdate is invalid with empty Details",
			fields: fields{
				Details: nil,
			},
			wantErr: true,
		},
		{
			name: "RecipeUpdate is invalid with Details slice of length 0",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailUpdate{},
			},
			wantErr: true,
		},
		{
			name: "RecipeUpdate is invalid with invalid detail",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailUpdate{
					{
						IngredientId: "",
						AmountNeed:   10.0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "RecipeUpdate is invalid with duplicate IngredientId",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailUpdate{
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
			name: "RecipeUpdate is valid",
			fields: fields{
				Details: []recipedetailmodel.RecipeDetailUpdate{
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
			data := &RecipeUpdate{
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
