package productmodel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductUpdateStatus_Validate(t *testing.T) {
	type fields struct {
		ProductId string
		IsActive  *bool
	}
	active := true
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ProductUpdateStatus is invalid with invalid product ID",
			fields: fields{
				ProductId: "",
				IsActive:  &active,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateStatus is invalid ils with nil isActive",
			fields: fields{
				ProductId: "Product123",
				IsActive:  nil,
			},
			wantErr: true,
		},
		{
			name: "ProductUpdateStatus is valid with valid data",
			fields: fields{
				ProductId: "Product123",
				IsActive:  &active,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ProductUpdateStatus{
				ProductId: tt.fields.ProductId,
				IsActive:  tt.fields.IsActive,
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
