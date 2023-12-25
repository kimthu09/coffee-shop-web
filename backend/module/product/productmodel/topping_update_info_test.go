package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/recipe/recipemodel"
	"coffee_shop_management_backend/module/recipedetail/recipedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToppingUpdate_TableName(t *testing.T) {
	type fields struct {
		ProductUpdateInfo *ProductUpdateInfo
		Cost              *int
		Price             *int
		Recipe            *recipemodel.RecipeUpdate
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of ToppingCreate successfully",
			fields: fields{
				ProductUpdateInfo: nil,
				Cost:              nil,
				Price:             nil,
				Recipe:            nil,
			},
			want: common.TableTopping,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topping := &ToppingUpdateInfo{
				ProductUpdateInfo: tt.fields.ProductUpdateInfo,
				Cost:              tt.fields.Cost,
				Price:             tt.fields.Price,
				Recipe:            tt.fields.Recipe,
			}
			got := topping.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestToppingUpdate_Validate(t *testing.T) {
	type fields struct {
		ProductUpdateInfo *ProductUpdateInfo
		Cost              *int
		Price             *int
		Recipe            *recipemodel.RecipeUpdate
	}

	str := "string"
	emptyString := ""
	number := 10
	negativeNumber := -10

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ToppingUpdate is valid with valid data",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name:         &str,
					Description:  &str,
					CookingGuide: &str,
				},
				Cost:  &number,
				Price: &number,
				Recipe: &recipemodel.RecipeUpdate{
					Details: []recipedetailmodel.RecipeDetailUpdate{
						{
							IngredientId: "Ing001",
							AmountNeed:   10,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ToppingUpdate is valid with nil data",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name:         nil,
					Description:  nil,
					CookingGuide: nil,
				},
				Cost:   nil,
				Price:  nil,
				Recipe: nil,
			},
			wantErr: false,
		},
		{
			name: "ToppingUpdate is invalid with invalid product information",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name:         &emptyString,
					Description:  nil,
					CookingGuide: nil,
				},
				Cost:   nil,
				Price:  nil,
				Recipe: nil,
			},
			wantErr: true,
		},
		{
			name: "ToppingUpdate is invalid with invalid cost",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name:         nil,
					Description:  nil,
					CookingGuide: nil,
				},
				Cost:   &negativeNumber,
				Price:  nil,
				Recipe: nil,
			},
			wantErr: true,
		},
		{
			name: "ToppingUpdate is invalid with invalid price",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name:         nil,
					Description:  nil,
					CookingGuide: nil,
				},
				Cost:   nil,
				Price:  &negativeNumber,
				Recipe: nil,
			},
			wantErr: true,
		},
		{
			name: "ToppingUpdate is invalid with invalid details",
			fields: fields{
				ProductUpdateInfo: &ProductUpdateInfo{
					Name:         nil,
					Description:  nil,
					CookingGuide: nil,
				},
				Cost:  nil,
				Price: nil,
				Recipe: &recipemodel.RecipeUpdate{
					Details: nil,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &ToppingUpdateInfo{
				ProductUpdateInfo: tt.fields.ProductUpdateInfo,
				Cost:              tt.fields.Cost,
				Price:             tt.fields.Price,
				Recipe:            tt.fields.Recipe,
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
