package exportnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportNoteCreate_TableName(t *testing.T) {
	type fields struct {
		Id                *string
		TotalPrice        float32
		CreateBy          string
		ExportNoteDetails []exportnotedetailmodel.ExportNoteDetailCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ExportNoteCreate successfully",
			fields: fields{
				Id:                nil,
				TotalPrice:        0,
				CreateBy:          "",
				ExportNoteDetails: nil,
			},
			want: common.TableExportNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exportNote := &ExportNoteCreate{
				Id:                tt.fields.Id,
				TotalPrice:        tt.fields.TotalPrice,
				CreateBy:          tt.fields.CreateBy,
				ExportNoteDetails: tt.fields.ExportNoteDetails,
			}

			got := exportNote.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestExportNoteCreate_Validate(t *testing.T) {
	type fields struct {
		Id                *string
		TotalPrice        float32
		CreateBy          string
		ExportNoteDetails []exportnotedetailmodel.ExportNoteDetailCreate
	}

	validExportNoteDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			ExportNoteId: "",
			IngredientId: "IngId1",
			ExpiryDate:   "01/02/2011",
			AmountExport: 12,
		},
	}
	invalidExportNoteDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			ExportNoteId: "",
			IngredientId: "IngId1",
			ExpiryDate:   "01/02/201313121",
			AmountExport: 0,
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
			name: "ExportNoteCreate is valid with nil id and valid CancelNoteDetail",
			fields: fields{
				Id:                nil,
				TotalPrice:        0,
				CreateBy:          "",
				ExportNoteDetails: validExportNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "ExportNoteCreate is valid with empty id and valid CancelNoteDetail",
			fields: fields{
				Id:                &emptyId,
				TotalPrice:        0,
				CreateBy:          "",
				ExportNoteDetails: validExportNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "ExportNoteCreate is invalid with invalidId and valid CancelNoteDetail",
			fields: fields{
				Id:                &invalidId,
				TotalPrice:        0,
				CreateBy:          "",
				ExportNoteDetails: validExportNoteDetails,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with valid Id and nil CancelNoteDetail",
			fields: fields{
				Id:                nil,
				TotalPrice:        0,
				CreateBy:          "",
				ExportNoteDetails: nil,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with valid Id and empty CancelNoteDetail",
			fields: fields{
				Id:         nil,
				TotalPrice: 0,
				CreateBy:   "",
				ExportNoteDetails: make(
					[]exportnotedetailmodel.ExportNoteDetailCreate, 0),
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with valid Id and invalid element inside CancelNoteDetail",
			fields: fields{
				Id:                nil,
				TotalPrice:        0,
				CreateBy:          "",
				ExportNoteDetails: invalidExportNoteDetails,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ExportNoteCreate{
				Id:                tt.fields.Id,
				TotalPrice:        tt.fields.TotalPrice,
				CreateBy:          tt.fields.CreateBy,
				ExportNoteDetails: tt.fields.ExportNoteDetails,
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
