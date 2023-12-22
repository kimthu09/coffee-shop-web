package exportnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportNoteDetailCreate_TableName(t *testing.T) {
	type fields struct {
		ExportNoteId string
		IngredientId string
		AmountExport int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ExportNoteDetailCreate successfully",
			fields: fields{
				ExportNoteId: "",
				IngredientId: "",
				AmountExport: 0,
			},
			want: common.TableExportNoteDetail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			detail := &ExportNoteDetailCreate{
				ExportNoteId: tt.fields.ExportNoteId,
				IngredientId: tt.fields.IngredientId,
				AmountExport: tt.fields.AmountExport,
			}
			got := detail.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestExportNoteDetailCreate_Validate(t *testing.T) {
	type fields struct {
		ExportNoteId string
		IngredientId string
		AmountExport int
	}

	mockExportNoteDetail := ExportNoteDetailCreate{
		IngredientId: "123456789",
		AmountExport: 1,
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ExportNoteDetailCreate is valid",
			fields: fields{
				IngredientId: mockExportNoteDetail.IngredientId,
				AmountExport: mockExportNoteDetail.AmountExport,
			},
			wantErr: false,
		},
		{
			name: "ExportNoteDetailCreate is invalid with invalid ingredient id",
			fields: fields{
				IngredientId: "",
				AmountExport: mockExportNoteDetail.AmountExport,
			},
			wantErr: true,
		},
		{
			name: "ExportNoteDetailCreate is invalid with not positive number",
			fields: fields{
				IngredientId: mockExportNoteDetail.IngredientId,
				AmountExport: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ExportNoteDetailCreate{
				ExportNoteId: tt.fields.ExportNoteId,
				IngredientId: tt.fields.IngredientId,
				AmountExport: tt.fields.AmountExport,
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
