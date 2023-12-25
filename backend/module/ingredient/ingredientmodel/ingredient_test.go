package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredient_TableName(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		Amount      int
		MeasureType *enum.MeasureType
		Price       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of Ingredient successfully",
			fields: fields{
				Id:          "",
				Name:        "",
				Amount:      0,
				MeasureType: nil,
				Price:       0,
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &Ingredient{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Amount:      tt.fields.Amount,
				MeasureType: tt.fields.MeasureType,
				Price:       tt.fields.Price,
			}
			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}
