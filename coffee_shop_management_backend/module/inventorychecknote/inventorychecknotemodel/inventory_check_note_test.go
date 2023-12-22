package inventorychecknotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInventoryCheckNote_TableName(t *testing.T) {
	type fields struct {
		Id                string
		AmountDifferent   int
		AmountAfterAdjust int
		CreatedBy         string
		CreatedByUser     usermodel.SimpleUser
		CreatedAt         *time.Time
		Details           []inventorychecknotedetailmodel.InventoryCheckNoteDetail
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of InventoryCheckNoteCreate successfully",
			fields: fields{
				Id:                "",
				AmountDifferent:   0,
				AmountAfterAdjust: 0,
				CreatedBy:         "",
				CreatedByUser:     usermodel.SimpleUser{},
				CreatedAt:         nil,
				Details:           nil,
			},
			want: common.TableInventoryCheckNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inventoryCheckNote := &InventoryCheckNote{
				Id:                tt.fields.Id,
				AmountDifferent:   tt.fields.AmountDifferent,
				AmountAfterAdjust: tt.fields.AmountAfterAdjust,
				CreatedBy:         tt.fields.CreatedBy,
				CreatedByUser:     tt.fields.CreatedByUser,
				CreatedAt:         tt.fields.CreatedAt,
				Details:           tt.fields.Details,
			}

			got := inventoryCheckNote.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}
