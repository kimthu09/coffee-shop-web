package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientUpdateAmount_TableName(t *testing.T) {
	type fields struct {
		Amount int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of IngredientUpdateAmount successfully",
			fields: fields{
				Amount: 0,
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &IngredientUpdateAmount{
				Amount: tt.fields.Amount,
			}
			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}
