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
		CreatedBy         string
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
				CreatedBy:         "",
				ExportNoteDetails: nil,
			},
			want: common.TableExportNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exportNote := &ExportNoteCreate{
				Id:                tt.fields.Id,
				CreatedBy:         tt.fields.CreatedBy,
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
		CreatedBy         string
		Reason            *ExportReason
		ExportNoteDetails []exportnotedetailmodel.ExportNoteDetailCreate
	}

	validExportNoteDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			IngredientId: "IngId1",
			AmountExport: 12,
		},
	}
	duplicateDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			IngredientId: "IngId1",
			AmountExport: 12,
		},
		{
			IngredientId: "IngId1",
			AmountExport: 11,
		},
	}
	invalidExportNoteDetails := []exportnotedetailmodel.ExportNoteDetailCreate{
		{
			IngredientId: "IngId1",
			AmountExport: 0,
		},
	}
	emptyId := ""
	invalidId := "This is invalid Id"
	reason := OutOfDate

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ExportNoteCreate is valid with id not exist",
			fields: fields{
				Id:                nil,
				Reason:            &reason,
				ExportNoteDetails: validExportNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "ExportNoteCreate is valid with empty id",
			fields: fields{
				Id:                &emptyId,
				Reason:            &reason,
				ExportNoteDetails: validExportNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "ExportNoteCreate is invalid with invalidId",
			fields: fields{
				Id:                &invalidId,
				Reason:            &reason,
				ExportNoteDetails: validExportNoteDetails,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with details not exist",
			fields: fields{
				Id:                nil,
				Reason:            &reason,
				ExportNoteDetails: nil,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with empty details",
			fields: fields{
				Id:     nil,
				Reason: &reason,
				ExportNoteDetails: make(
					[]exportnotedetailmodel.ExportNoteDetailCreate, 0),
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with invalid details",
			fields: fields{
				Id:                nil,
				Reason:            &reason,
				ExportNoteDetails: invalidExportNoteDetails,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with duplicate detail",
			fields: fields{
				Id:                nil,
				Reason:            &reason,
				ExportNoteDetails: duplicateDetails,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteCreate is invalid with reason not exist",
			fields: fields{
				Id:                nil,
				Reason:            nil,
				ExportNoteDetails: invalidExportNoteDetails,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ExportNoteCreate{
				Id:                tt.fields.Id,
				CreatedBy:         tt.fields.CreatedBy,
				Reason:            tt.fields.Reason,
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
