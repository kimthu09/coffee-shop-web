package recipedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecipeDetailCreate_TableName(t *testing.T) {
	type fields struct {
		RecipeId     string
		IngredientId string
		AmountNeed   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of RecipeDetailCreate successfully",
			fields: fields{
				RecipeId:     "",
				IngredientId: "",
				AmountNeed:   0,
			},
			want: common.TableRecipeDetail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recipeDetail := &RecipeDetailCreate{
				RecipeId:     tt.fields.RecipeId,
				IngredientId: tt.fields.IngredientId,
				AmountNeed:   tt.fields.AmountNeed,
			}

			got := recipeDetail.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestRecipeDetailCreate_Validate(t *testing.T) {
	type fields struct {
		RecipeId     string
		IngredientId string
		AmountNeed   int
	}

	recipeId := "recipe123"
	invalidId := "123456789001234567890"
	validId := "0123456789"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "RecipeDetailCreate is invalid with invalid IngredientId",
			fields: fields{
				RecipeId:     recipeId,
				IngredientId: invalidId,
				AmountNeed:   10.0,
			},
			wantErr: true,
		},
		{
			name: "RecipeDetailCreate is invalid with not positive AmountNeed",
			fields: fields{
				RecipeId:     recipeId,
				IngredientId: validId,
				AmountNeed:   0,
			},
			wantErr: true,
		},
		{
			name: "RecipeDetailCreate is valid",
			fields: fields{
				RecipeId:     recipeId,
				IngredientId: validId,
				AmountNeed:   10.0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &RecipeDetailCreate{
				RecipeId:     tt.fields.RecipeId,
				IngredientId: tt.fields.IngredientId,
				AmountNeed:   tt.fields.AmountNeed,
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
