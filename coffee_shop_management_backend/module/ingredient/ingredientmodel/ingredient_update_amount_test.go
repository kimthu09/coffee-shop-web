package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientUpdateAmount_TableName(t *testing.T) {
	type fields struct {
		Amount float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of IngredientUpdateAmount successfully",
			fields: fields{
				Amount: 0,
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &IngredientUpdateAmount{
				Amount: tt.fields.Amount,
			}
			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestIngredientUpdateAmount_Validate(t *testing.T) {
	type fields struct {
		Amount float32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "IngredientUpdateAmount is invalid with amount equal 0",
			fields: fields{
				Amount: 0,
			},
			wantErr: true,
		},
		{
			name: "IngredientUpdateAmount is valid",
			fields: fields{
				Amount: -1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &IngredientUpdateAmount{
				Amount: tt.fields.Amount,
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
