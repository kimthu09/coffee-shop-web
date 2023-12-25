package categoryfoodmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/category/categorymodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryFood_TableName(t *testing.T) {
	type fields struct {
		FoodId     string
		CategoryId string
		Category   categorymodel.Category
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CategoryFood successfully",
			fields: fields{
				FoodId:     "",
				CategoryId: "",
				Category:   categorymodel.Category{},
			},
			want: common.TableCategoryFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categoryFood := &CategoryFood{
				FoodId:     tt.fields.FoodId,
				CategoryId: tt.fields.CategoryId,
				Category:   tt.fields.Category,
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
