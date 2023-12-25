package inventorychecknotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInventoryCheckNoteCreate_TableName(t *testing.T) {
	type fields struct {
		Id                *string
		AmountDifferent   int
		AmountAfterAdjust int
		CreatedBy         string
		Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of InventoryCheckNoteCreate successfully",
			fields: fields{
				Id:                nil,
				AmountDifferent:   0,
				AmountAfterAdjust: 0,
				CreatedBy:         "",
				Details:           nil,
			},
			want: common.TableInventoryCheckNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inventoryCheckNote := &InventoryCheckNoteCreate{
				Id:                tt.fields.Id,
				AmountDifferent:   tt.fields.AmountDifferent,
				AmountAfterAdjust: tt.fields.AmountAfterAdjust,
				CreatedBy:         tt.fields.CreatedBy,
				Details:           tt.fields.Details,
			}

			got := inventoryCheckNote.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestInventoryCheckNoteCreate_Validate(t *testing.T) {
	type fields struct {
		Id                *string
		AmountDifferent   int
		AmountAfterAdjust int
		CreatedBy         string
		Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate
	}
	invalidId := "12356012356890"
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "InventoryCheckNoteCreate is invalid with invalid id",
			fields: fields{
				Id:      &invalidId,
				Details: nil,
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteCreate is invalid with details not exist",
			fields: fields{
				Id:      nil,
				Details: nil,
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteCreate is invalid with empty details",
			fields: fields{
				Id:      nil,
				Details: make([]inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate, 0),
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteCreate is invalid with invalid detail",
			fields: fields{
				Id: nil,
				Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
					{
						IngredientId: "Ingredient1",
						Difference:   0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteCreate is invalid with duplicate ingredient",
			fields: fields{
				Id: nil,
				Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
					{
						IngredientId: "Ingredient1",
						Difference:   100,
					},
					{
						IngredientId: "Ingredient1",
						Difference:   200,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "InventoryCheckNoteCreate is valid",
			fields: fields{
				Id: nil,
				Details: []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
					{
						IngredientId: "Ingredient1",
						Difference:   100,
					},
					{
						IngredientId: "Ingredient2",
						Difference:   200,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &InventoryCheckNoteCreate{
				Id:                tt.fields.Id,
				AmountDifferent:   tt.fields.AmountDifferent,
				AmountAfterAdjust: tt.fields.AmountAfterAdjust,
				CreatedBy:         tt.fields.CreatedBy,
				Details:           tt.fields.Details,
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
