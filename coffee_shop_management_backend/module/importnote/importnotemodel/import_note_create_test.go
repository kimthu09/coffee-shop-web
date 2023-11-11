package importnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImportNoteCreate_TableName(t *testing.T) {
	type fields struct {
		Id                *string
		TotalPrice        float32
		SupplierId        string
		CreateBy          string
		ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ImportNoteCreate successfully",
			fields: fields{
				Id:                nil,
				TotalPrice:        0,
				SupplierId:        "",
				CreateBy:          "",
				ImportNoteDetails: nil,
			},
			want: common.TableImportNote,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			importNote := &ImportNoteCreate{
				Id:                tt.fields.Id,
				TotalPrice:        tt.fields.TotalPrice,
				SupplierId:        tt.fields.SupplierId,
				CreateBy:          tt.fields.CreateBy,
				ImportNoteDetails: tt.fields.ImportNoteDetails,
			}

			got := importNote.TableName()

			assert.Equal(
				t, tt.want, got,
				"TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestImportNoteCreate_Validate(t *testing.T) {
	type fields struct {
		Id                *string
		TotalPrice        float32
		SupplierId        string
		CreateBy          string
		ImportNoteDetails []importnotedetailmodel.ImportNoteDetailCreate
	}

	validId := "101234567890"
	invalidId := "1312321321312"
	validImportNoteDetail := importnotedetailmodel.ImportNoteDetailCreate{
		IngredientId:   validId,
		ExpiryDate:     "09/11/2023",
		Price:          10000,
		IsReplacePrice: true,
		AmountImport:   100,
	}
	duplicateImportNoteDetail := importnotedetailmodel.ImportNoteDetailCreate{
		IngredientId:   validId,
		ExpiryDate:     "09/11/2023",
		Price:          12,
		IsReplacePrice: true,
		AmountImport:   12,
	}
	invalidImportNoteDetail := importnotedetailmodel.ImportNoteDetailCreate{
		IngredientId: "",
		ExpiryDate:   "09/11/2023",
		Price:        10000,
		AmountImport: 100,
	}
	validImportNote := ImportNoteCreate{
		Id:         &validId,
		SupplierId: validId,
		ImportNoteDetails: []importnotedetailmodel.ImportNoteDetailCreate{
			validImportNoteDetail,
		},
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ImportNoteCreate is valid",
			fields: fields{
				Id:                validImportNote.Id,
				TotalPrice:        validImportNote.TotalPrice,
				SupplierId:        validImportNote.SupplierId,
				CreateBy:          validImportNote.CreateBy,
				ImportNoteDetails: validImportNote.ImportNoteDetails,
			},
			wantErr: false,
		},
		{
			name: "ImportNoteCreate is invalid with invalid id",
			fields: fields{
				Id:                &invalidId,
				TotalPrice:        validImportNote.TotalPrice,
				SupplierId:        validImportNote.SupplierId,
				CreateBy:          validImportNote.CreateBy,
				ImportNoteDetails: validImportNote.ImportNoteDetails,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteCreate is invalid with invalid supplier id",
			fields: fields{
				Id:                &validId,
				TotalPrice:        validImportNote.TotalPrice,
				SupplierId:        "",
				CreateBy:          validImportNote.CreateBy,
				ImportNoteDetails: validImportNote.ImportNoteDetails,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteCreate is invalid with invalid ImportNoteDetails",
			fields: fields{
				Id:                &validId,
				TotalPrice:        validImportNote.TotalPrice,
				SupplierId:        validImportNote.SupplierId,
				CreateBy:          validImportNote.CreateBy,
				ImportNoteDetails: nil,
			},
			wantErr: true,
		},
		{
			name: "ImportNoteCreate is invalid with invalid element inside ImportNoteDetails",
			fields: fields{
				Id:         &validId,
				TotalPrice: validImportNote.TotalPrice,
				SupplierId: validImportNote.SupplierId,
				CreateBy:   validImportNote.CreateBy,
				ImportNoteDetails: []importnotedetailmodel.ImportNoteDetailCreate{
					invalidImportNoteDetail,
				},
			},
			wantErr: true,
		},
		{
			name: "ImportNoteCreate is invalid with duplicate ingredient detail need to replace price",
			fields: fields{
				Id:         &validId,
				TotalPrice: validImportNote.TotalPrice,
				SupplierId: validImportNote.SupplierId,
				CreateBy:   validImportNote.CreateBy,
				ImportNoteDetails: []importnotedetailmodel.ImportNoteDetailCreate{
					validImportNoteDetail,
					duplicateImportNoteDetail,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ImportNoteCreate{
				Id:                tt.fields.Id,
				TotalPrice:        tt.fields.TotalPrice,
				SupplierId:        tt.fields.SupplierId,
				CreateBy:          tt.fields.CreateBy,
				ImportNoteDetails: tt.fields.ImportNoteDetails,
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
