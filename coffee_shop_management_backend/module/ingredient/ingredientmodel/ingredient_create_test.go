package ingredientmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIngredientCreate_TableName(t *testing.T) {
	type fields struct {
		Id          *string
		Name        string
		MeasureType *enum.MeasureType
		Price       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of IngredientCreate successfully",
			fields: fields{
				Id:          nil,
				Name:        "",
				MeasureType: nil,
				Price:       0,
			},
			want: common.TableIngredient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ingredient := &IngredientCreate{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				MeasureType: tt.fields.MeasureType,
				Price:       tt.fields.Price,
			}
			got := ingredient.TableName()
			assert.Equal(t, tt.want, got, "TableName() = %v, want %v", got, tt.want)
		})
	}
}

func TestIngredientCreate_Validate(t *testing.T) {
	type fields struct {
		Id          *string
		Name        string
		MeasureType *enum.MeasureType
		Price       float32
	}

	invalidId := "012345678901ehehehe"
	measureType := enum.Weight
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "IngredientCreate is invalid with invalid id",
			fields: fields{
				Id: &invalidId,
			},
			wantErr: true,
		},
		{
			name: "IngredientCreate is invalid with empty name",
			fields: fields{
				Id:   nil,
				Name: "",
			},
			wantErr: true,
		},
		{
			name: "IngredientCreate is invalid with nil measure type",
			fields: fields{
				Id:          nil,
				Name:        "123",
				MeasureType: nil,
			},
			wantErr: true,
		},
		{
			name: "IngredientCreate is invalid with negative price",
			fields: fields{
				Id:          nil,
				Name:        "123",
				MeasureType: &measureType,
				Price:       -1,
			},
			wantErr: true,
		},
		{
			name: "IngredientCreate is valid",
			fields: fields{
				Id:          nil,
				Name:        "123",
				MeasureType: &measureType,
				Price:       0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &IngredientCreate{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				MeasureType: tt.fields.MeasureType,
				Price:       tt.fields.Price,
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
