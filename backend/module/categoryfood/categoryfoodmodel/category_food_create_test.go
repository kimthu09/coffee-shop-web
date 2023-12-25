package categoryfoodmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryFoodCreate_TableName(t *testing.T) {
	type fields struct {
		FoodId     string
		CategoryId string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CategoryFoodCreate successfully",
			fields: fields{
				FoodId:     "",
				CategoryId: "",
			},
			want: common.TableCategoryFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categoryFood := &CategoryFoodCreate{
				FoodId:     tt.fields.FoodId,
				CategoryId: tt.fields.CategoryId,
			}
			got := categoryFood.TableName()

			assert.Equal(
				t,
				tt.want,
				got,
				"TableName() = %v, want %v", got, tt.want)
		})
	}
}
