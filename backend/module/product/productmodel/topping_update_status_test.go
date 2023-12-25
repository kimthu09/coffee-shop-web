package productmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToppingUpdateStatus_TableName(t *testing.T) {
	type fields struct {
		ProductUpdateStatus *ProductUpdateStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ToppingUpdateStatus successfully",
			fields: fields{
				ProductUpdateStatus: nil,
			},
			want: common.TableTopping,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topping := &ToppingUpdateStatus{
				ProductUpdateStatus: tt.fields.ProductUpdateStatus,
			}
			got := topping.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestToppingUpdateStatus_Validate(t *testing.T) {
	type fields struct {
		ProductUpdateStatus *ProductUpdateStatus
	}

	active := true

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ToppingUpdateStatus is valid with valid product information",
			fields: fields{
				ProductUpdateStatus: &ProductUpdateStatus{
					ProductId: "Product123",
					IsActive:  &active,
				},
			},
			wantErr: false,
		},
		{
			name: "ToppingUpdateStatus is invalid with invalid product information",
			fields: fields{
				ProductUpdateStatus: &ProductUpdateStatus{
					ProductId: "",
					IsActive:  &active,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ToppingUpdateStatus{
				ProductUpdateStatus: tt.fields.ProductUpdateStatus,
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
