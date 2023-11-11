package suppliermodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplierUpdateDebt_TableName(t *testing.T) {
	type fields struct {
		Amount   *float32
		CreateBy string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of SupplierUpdateDebt successfully",
			fields: fields{
				Amount:   nil,
				CreateBy: "",
			},
			want: common.TableSupplier,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			supplier := &SupplierUpdateDebt{
				Amount:   tt.fields.Amount,
				CreateBy: tt.fields.CreateBy,
			}
			got := supplier.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestSupplierUpdateDebt_Validate(t *testing.T) {
	type fields struct {
		Amount   *float32
		CreateBy string
	}
	invalidAmount := float32(0)
	validAmount := float32(123)
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "SupplierUpdateDebt is invalid with nil amount",
			fields: fields{
				Amount: nil,
			},
			wantErr: true,
		},
		{
			name: "SupplierUpdateDebt is invalid with amount equal 0",
			fields: fields{
				Amount: &invalidAmount,
			},
			wantErr: true,
		},
		{
			name: "SupplierUpdateDebt is successfully",
			fields: fields{
				Amount: &validAmount,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &SupplierUpdateDebt{
				Amount:   tt.fields.Amount,
				CreateBy: tt.fields.CreateBy,
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
