package productmodel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductUpdateInfo_Validate(t *testing.T) {
	type fields struct {
		Name         *string
		Description  *string
		CookingGuide *string
		IsActive     *bool
	}

	name := "Updated Product Name"
	emptyName := ""
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ProductUpdateInfo is valid with valid data",
			fields: fields{
				Name: &name,
			},
			wantErr: false,
		},
		{
			name:    "ProductUpdateInfo is valid with name not exist",
			fields:  fields{},
			wantErr: false,
		},
		{
			name: "ProductUpdate is invalid with empty product name",
			fields: fields{
				Name: &emptyName,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ProductUpdateInfo{
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
