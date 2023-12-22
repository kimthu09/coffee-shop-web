package inventorychecknotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventoryCheckNoteDetailCreate_TableName(t *testing.T) {
	type fields struct {
		InventoryCheckNoteId string
		IngredientId         string
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
			name: "Get TableName of InventoryCheckNoteDetailCreate successfully",
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
			inventoryCheckNoteDetail := &InventoryCheckNoteDetailCreate{
				InventoryCheckNoteId: tt.fields.InventoryCheckNoteId,
				IngredientId:         tt.fields.IngredientId,
				Initial:              tt.fields.Initial,
				Difference:           tt.fields.Difference,
				Final:                tt.fields.Final,
			}

			got := inventoryCheckNoteDetail.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestInventoryCheckNoteDetailCreate_Validate(t *testing.T) {
	type fields struct {
		InventoryCheckNoteId string
		IngredientId         string
		Initial              int
		Difference           int
		Final                int
	}

	validId := "123567890"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "InventoryCheckNoteDetailCreate is invalid with invalid ingredient",
			fields: fields{
				IngredientId: "",
				Difference:   0,
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteDetailCreate is invalid with zero difference",
			fields: fields{
				IngredientId: validId,
				Difference:   0,
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteDetailCreate is valid",
			fields: fields{
				IngredientId: validId,
				Difference:   10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &InventoryCheckNoteDetailCreate{
				InventoryCheckNoteId: tt.fields.InventoryCheckNoteId,
				IngredientId:         tt.fields.IngredientId,
				Initial:              tt.fields.Initial,
				Difference:           tt.fields.Difference,
				Final:                tt.fields.Final,
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
