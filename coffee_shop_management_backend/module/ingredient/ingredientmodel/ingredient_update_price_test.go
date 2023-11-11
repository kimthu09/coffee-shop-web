package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientUpdatePrice_TableName(t *testing.T) {
	type fields struct {
		Price *float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of IngredientUpdatePrice successfully",
			fields: fields{
				Price: nil,
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &IngredientUpdatePrice{
				Price: tt.fields.Price,
			}
			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestIngredientUpdatePrice_Validate(t *testing.T) {
	type fields struct {
		Price *float32
	}

	price := float32(-1)
	notNegativePrice := float32(0)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "IngredientUpdatePrice is invalid with price negative",
			fields: fields{
				Price: &price,
			},
			wantErr: true,
		},
		{
			name: "IngredientUpdatePrice is valid with nil price",
			fields: fields{
				Price: nil,
			},
			wantErr: false,
		},
		{
			name: "IngredientUpdatePrice is valid with not negative price",
			fields: fields{
				Price: &notNegativePrice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &IngredientUpdatePrice{
				Price: tt.fields.Price,
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
