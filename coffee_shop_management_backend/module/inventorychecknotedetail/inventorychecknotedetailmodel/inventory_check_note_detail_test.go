package inventorychecknotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventoryCheckNoteDetail_TableName(t *testing.T) {
	type fields struct {
		InventoryCheckNoteId string
		IngredientId         string
		Ingredient           ingredientmodel.SimpleIngredient
		Initial              int
		Difference           int
		Final                int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of InventoryCheckNoteDetail successfully",
			fields: fields{
				InventoryCheckNoteId: "",
				IngredientId:         "",
				Initial:              0,
				Difference:           0,
				Final:                0,
			},
			want: common.TableInventoryCheckNoteDetail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inventoryCheckNoteDetail := &InventoryCheckNoteDetail{
				InventoryCheckNoteId: tt.fields.InventoryCheckNoteId,
				IngredientId:         tt.fields.IngredientId,
				Ingredient:           tt.fields.Ingredient,
				Initial:              tt.fields.Initial,
				Difference:           tt.fields.Difference,
				Final:                tt.fields.Final,
			}

			got := inventoryCheckNoteDetail.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}
