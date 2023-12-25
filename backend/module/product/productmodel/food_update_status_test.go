package productmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoodUpdateStatus_TableName(t *testing.T) {
	type fields struct {
		ProductUpdateStatus *ProductUpdateStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of FoodUpdateStatus successfully",
			fields: fields{
				ProductUpdateStatus: nil,
			},
			want: common.TableFood,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			food := &FoodUpdateStatus{
				ProductUpdateStatus: tt.fields.ProductUpdateStatus,
			}
			got := food.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestFoodUpdateStatus_Validate(t *testing.T) {
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
			name: "FoodUpdateStatus is valid with valid product information",
			fields: fields{
				ProductUpdateStatus: &ProductUpdateStatus{
					ProductId: "Product123",
					IsActive:  &active,
				},
			},
			wantErr: false,
		},
		{
			name: "FoodUpdateStatus is invalid with invalid product information",
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
			data := &FoodUpdateStatus{
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
