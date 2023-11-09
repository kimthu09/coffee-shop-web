package cancelnotedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCancelNoteDetailCreate_TableName(t *testing.T) {
	type fields struct {
		CancelNoteId string
		IngredientId string
		ExpiryDate   string
		Reason       *CancelReason
		AmountCancel float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of CancelNoteDetailCreate successfully",
			fields: fields{
				CancelNoteId: "",
				IngredientId: "",
				ExpiryDate:   "",
				Reason:       nil,
				AmountCancel: 0,
			},
			want: common.TableCancelNoteDetail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			detail := &CancelNoteDetailCreate{
				CancelNoteId: tt.fields.CancelNoteId,
				IngredientId: tt.fields.IngredientId,
				ExpiryDate:   tt.fields.ExpiryDate,
				Reason:       tt.fields.Reason,
				AmountCancel: tt.fields.AmountCancel,
			}
			got := detail.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestCancelNoteDetailCreate_Validate(t *testing.T) {
	type fields struct {
		CancelNoteId string
		IngredientId string
		ExpiryDate   string
		Reason       *CancelReason
		AmountCancel float32
	}

	mockReason := Damaged
	mockCancelNoteDetail := CancelNoteDetailCreate{
		IngredientId: "123456789",
		ExpiryDate:   "08/11/2023",
		Reason:       &mockReason,
		AmountCancel: 1,
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "CancelNoteDetailCreate is valid",
			fields: fields{
				IngredientId: mockCancelNoteDetail.IngredientId,
				ExpiryDate:   mockCancelNoteDetail.ExpiryDate,
				Reason:       mockCancelNoteDetail.Reason,
				AmountCancel: mockCancelNoteDetail.AmountCancel,
			},
			wantErr: false,
		},
		{
			name: "CancelNoteDetailCreate is invalid with invalid ingredient id",
			fields: fields{
				IngredientId: "",
				ExpiryDate:   mockCancelNoteDetail.ExpiryDate,
				Reason:       mockCancelNoteDetail.Reason,
				AmountCancel: mockCancelNoteDetail.AmountCancel,
			},
			wantErr: true,
		},
		{
			name: "CancelNoteDetailCreate is invalid with invalid expiry date",
			fields: fields{
				IngredientId: mockCancelNoteDetail.IngredientId,
				ExpiryDate:   "This is invalid expiry date",
				Reason:       mockCancelNoteDetail.Reason,
				AmountCancel: mockCancelNoteDetail.AmountCancel,
			},
			wantErr: true,
		},
		{
			name: "CancelNoteDetailCreate is invalid with nil reason",
			fields: fields{
				IngredientId: mockCancelNoteDetail.IngredientId,
				ExpiryDate:   mockCancelNoteDetail.ExpiryDate,
				Reason:       nil,
				AmountCancel: mockCancelNoteDetail.AmountCancel,
			},
			wantErr: true,
		},
		{
			name: "CancelNoteDetailCreate is invalid with not positive number",
			fields: fields{
				IngredientId: mockCancelNoteDetail.IngredientId,
				ExpiryDate:   mockCancelNoteDetail.ExpiryDate,
				Reason:       mockCancelNoteDetail.Reason,
				AmountCancel: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &CancelNoteDetailCreate{
				CancelNoteId: tt.fields.CancelNoteId,
				IngredientId: tt.fields.IngredientId,
				ExpiryDate:   tt.fields.ExpiryDate,
				Reason:       tt.fields.Reason,
				AmountCancel: tt.fields.AmountCancel,
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
