package suppliermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplierUpdateDebt_TableName(t *testing.T) {
	type fields struct {
		Id        *string
		Amount    *int
		CreatedBy string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of SupplierUpdateDebt successfully",
			fields: fields{
				Id:        nil,
				Amount:    nil,
				CreatedBy: "",
			},
			want: common.TableSupplier,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			supplier := &SupplierUpdateDebt{
				Id:        tt.fields.Id,
				Amount:    tt.fields.Amount,
				CreatedBy: tt.fields.CreatedBy,
			}
			got := supplier.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestSupplierUpdateDebt_Validate(t *testing.T) {
	type fields struct {
		Id        *string
		Amount    *int
		CreatedBy string
	}
	invalidId := "0123456789012345678901234567890"
	invalidAmount := 0
	validAmount := 123
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "SupplierUpdateDebt is invalid with invalid id",
			fields: fields{
				Id:     &invalidId,
				Amount: &validAmount,
			},
			wantErr: true,
		},
		{
			name: "SupplierUpdateDebt is invalid with nil amount",
			fields: fields{
				Id:     nil,
				Amount: nil,
			},
			wantErr: true,
		},
		{
			name: "SupplierUpdateDebt is invalid with amount equal 0",
			fields: fields{
				Id:     nil,
				Amount: &invalidAmount,
			},
			wantErr: true,
		},
		{
			name: "SupplierUpdateDebt is successfully",
			fields: fields{
				Id:     nil,
				Amount: &validAmount,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &SupplierUpdateDebt{
				Id:        tt.fields.Id,
				Amount:    tt.fields.Amount,
				CreatedBy: tt.fields.CreatedBy,
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
