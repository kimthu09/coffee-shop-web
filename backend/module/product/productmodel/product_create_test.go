package productmodel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductCreate_Validate(t *testing.T) {
	type fields struct {
		Id           *string
		Name         string
		Description  string
		CookingGuide string
	}

	validId := "product123"
	invalidId := "12345678901234567890"

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ProductCreate is valid with valid data",
			fields: fields{
				Id:           &validId,
				Name:         "Product Name",
				Description:  "Product Description",
				CookingGuide: "Cooking Guide",
			},
			wantErr: false,
		},
		{
			name: "ProductCreate is invalid with invalid product Id",
			fields: fields{
				Id: &invalidId,
			},
			wantErr: true,
		},
		{
			name: "ProductCreate is invalid with empty product name",
			fields: fields{
				Id:   &validId,
				Name: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ProductCreate{
				Id:           tt.fields.Id,
				Name:         tt.fields.Name,
				Description:  tt.fields.Description,
				CookingGuide: tt.fields.CookingGuide,
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
