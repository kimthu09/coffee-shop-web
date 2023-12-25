package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleIngredient_TableName(t *testing.T) {
	type fields struct {
		Id          string
		Name        string
		MeasureType *enum.MeasureType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of SimpleIngredient successfully",
			fields: fields{
				Name:        "",
				MeasureType: nil,
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &SimpleIngredient{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				MeasureType: tt.fields.MeasureType,
			}

			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}
