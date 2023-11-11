package cancelnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelNoteCreate_TableName(t *testing.T) {
	type fields struct {
		Id                      *string
		TotalPrice              float32
		CreateBy                string
		CancelNoteCreateDetails []cancelnotedetailmodel.CancelNoteDetailCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CancelNoteCreate successfully",
			fields: fields{
				Id:                      nil,
				TotalPrice:              0,
				CreateBy:                "",
				CancelNoteCreateDetails: nil,
			},
			want: common.TableCancelNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cancelNote := &CancelNoteCreate{
				Id:                      tt.fields.Id,
				TotalPrice:              tt.fields.TotalPrice,
				CreateBy:                tt.fields.CreateBy,
				CancelNoteCreateDetails: tt.fields.CancelNoteCreateDetails,
			}
			got := cancelNote.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestCancelNoteCreate_Validate(t *testing.T) {
	type fields struct {
		Id                      *string
		TotalPrice              float32
		CreateBy                string
		CancelNoteCreateDetails []cancelnotedetailmodel.CancelNoteDetailCreate
	}

	reason := cancelnotedetailmodel.Damaged
	validCancelNoteDetails := []cancelnotedetailmodel.CancelNoteDetailCreate{
		{
			CancelNoteId: "",
			IngredientId: "IngId1",
			ExpiryDate:   "01/02/2011",
			Reason:       &reason,
			AmountCancel: 12,
		},
	}
	invalidCancelNoteDetails := []cancelnotedetailmodel.CancelNoteDetailCreate{
		{
			CancelNoteId: "",
			IngredientId: "IngId1",
			ExpiryDate:   "01/02/201313121",
			Reason:       &reason,
			AmountCancel: 0,
		},
	}
	emptyId := ""
	invalidId := "This is invalid Id"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "CancelNoteCreate is valid with nil id and valid CancelNoteDetail",
			fields: fields{
				Id:                      nil,
				TotalPrice:              0,
				CreateBy:                "",
				CancelNoteCreateDetails: validCancelNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "CancelNoteCreate is valid with empty id and valid CancelNoteDetail",
			fields: fields{
				Id:                      &emptyId,
				TotalPrice:              0,
				CreateBy:                "",
				CancelNoteCreateDetails: validCancelNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "CancelNoteCreate is invalid with invalidId and valid CancelNoteDetail",
			fields: fields{
				Id:                      &invalidId,
				TotalPrice:              0,
				CreateBy:                "",
				CancelNoteCreateDetails: validCancelNoteDetails,
			},
			wantErr: true,
		},
		{
			name: "CancelNoteCreate is invalid with valid Id and nil CancelNoteDetail",
			fields: fields{
				Id:                      nil,
				TotalPrice:              0,
				CreateBy:                "",
				CancelNoteCreateDetails: nil,
			},
			wantErr: true,
		},
		{
			name: "CancelNoteCreate is invalid with valid Id and empty CancelNoteDetail",
			fields: fields{
				Id:         nil,
				TotalPrice: 0,
				CreateBy:   "",
				CancelNoteCreateDetails: make(
					[]cancelnotedetailmodel.CancelNoteDetailCreate, 0),
			},
			wantErr: true,
		},
		{
			name: "CancelNoteCreate is invalid with valid Id and invalid element inside CancelNoteDetail",
			fields: fields{
				Id:                      nil,
				TotalPrice:              0,
				CreateBy:                "",
				CancelNoteCreateDetails: invalidCancelNoteDetails,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &CancelNoteCreate{
				Id:                      tt.fields.Id,
				TotalPrice:              tt.fields.TotalPrice,
				CreateBy:                tt.fields.CreateBy,
				CancelNoteCreateDetails: tt.fields.CancelNoteCreateDetails,
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
