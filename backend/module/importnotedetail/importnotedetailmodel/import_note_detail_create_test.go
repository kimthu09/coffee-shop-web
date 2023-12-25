package importnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportNoteDetailCreate_TableName(t *testing.T) {
	type fields struct {
		ImportNoteId   string
		IngredientId   string
		ExpiryDate     string
		Price          float32
		IsReplacePrice bool
		AmountImport   int
		TotalUnit      float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ImportNoteDetailCreate successfully",
			fields: fields{
				ImportNoteId:   "",
				IngredientId:   "",
				ExpiryDate:     "",
				Price:          0,
				IsReplacePrice: false,
				AmountImport:   0,
			},
			want: common.TableImportNoteDetail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			importNoteDetail := &ImportNoteDetailCreate{
				ImportNoteId:   tt.fields.ImportNoteId,
				IngredientId:   tt.fields.IngredientId,
				Price:          tt.fields.Price,
				IsReplacePrice: tt.fields.IsReplacePrice,
				AmountImport:   tt.fields.AmountImport,
			}
			got := importNoteDetail.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestImportNoteDetailCreate_Validate(t *testing.T) {
	type fields struct {
		ImportNoteId   string
		IngredientId   string
		Price          float32
		IsReplacePrice bool
		AmountImport   int
		TotalUnit      float32
	}
	validId := "012345678901"
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ImportNoteDetailCreate is invalid with invalid id ingredient",
			fields: fields{
				IngredientId: "",
			},
			wantErr: true,
		},
		{
			name: "ImportNoteDetailCreate is invalid with negative price",
			fields: fields{
				IngredientId: validId,
				Price:        -1,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteDetailCreate is invalid with not positive amount import",
			fields: fields{
				IngredientId: validId,
				Price:        0,
				AmountImport: 0,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteDetailCreate is valid",
			fields: fields{
				IngredientId: validId,
				Price:        0,
				AmountImport: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ImportNoteDetailCreate{
				ImportNoteId:   tt.fields.ImportNoteId,
				IngredientId:   tt.fields.IngredientId,
				Price:          tt.fields.Price,
				IsReplacePrice: tt.fields.IsReplacePrice,
				AmountImport:   tt.fields.AmountImport,
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
