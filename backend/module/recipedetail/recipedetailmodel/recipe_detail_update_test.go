package recipedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecipeDetailUpdate_TableName(t *testing.T) {
	type fields struct {
		IngredientId string
		AmountNeed   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of RecipeDetailUpdate successfully",
			fields: fields{
				IngredientId: "",
				AmountNeed:   0,
			},
			want: common.TableRecipeDetail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recipeDetail := &RecipeDetailUpdate{
				IngredientId: tt.fields.IngredientId,
				AmountNeed:   tt.fields.AmountNeed,
			}

			got := recipeDetail.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestRecipeDetailUpdate_Validate(t *testing.T) {
	type fields struct {
		IngredientId string
		AmountNeed   int
	}

	invalidId := "123456789001234567890"
	validId := "0123456789"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "RecipeDetailUpdate is invalid with invalid IngredientId",
			fields: fields{
				IngredientId: invalidId,
				AmountNeed:   10.0,
			},
			wantErr: true,
		},
		{
			name: "RecipeDetailUpdate is invalid with not positive AmountNeed",
			fields: fields{
				IngredientId: validId,
				AmountNeed:   -5.0,
			},
			wantErr: true,
		},
		{
			name: "RecipeDetailUpdate is valid",
			fields: fields{
				IngredientId: validId,
				AmountNeed:   10.0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &RecipeDetailUpdate{
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
