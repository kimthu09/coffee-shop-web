package invoicedetailmodel

import (
	"coffee_shop_management_backend/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceDetailCreate_TableName(t *testing.T) {
	type fields struct {
		InvoiceId   string
		FoodId      string
		FoodName    string
		SizeId      string
		SizeName    string
		Toppings    *InvoiceDetailToppings
		Amount      int
		UnitPrice   int
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of InvoiceDetailCreate successfully",
			fields: fields{
				InvoiceId:   "",
				FoodId:      "",
				FoodName:    "",
				SizeId:      "",
				SizeName:    "",
				Toppings:    nil,
				Amount:      0,
				UnitPrice:   0,
				Description: "",
			},
			want: common.TableInvoiceDetail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoice := &InvoiceDetailCreate{
				InvoiceId:   tt.fields.InvoiceId,
				FoodId:      tt.fields.FoodId,
				FoodName:    tt.fields.FoodName,
				SizeId:      tt.fields.SizeId,
				SizeName:    tt.fields.SizeName,
				Toppings:    tt.fields.Toppings,
				Amount:      tt.fields.Amount,
				UnitPrice:   tt.fields.UnitPrice,
				Description: tt.fields.Description,
			}
			got := invoice.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)

		})
	}
}

func TestInvoiceDetailCreate_Validate(t *testing.T) {
	type fields struct {
		InvoiceId   string
		FoodId      string
		FoodName    string
		SizeId      string
		SizeName    string
		Toppings    *InvoiceDetailToppings
		Amount      int
		UnitPrice   int
		Description string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "InvoiceDetailCreate is invalid with nil FoodId",
			fields: fields{
				FoodId: "",
			},
			wantErr: true,
		},
		{
			name: "InvoiceDetailCreate is invalid with nil SizeId",
			fields: fields{
				FoodId: "food123",
				SizeId: "",
			},
			wantErr: true,
		},
		{
			name: "InvoiceDetailCreate is invalid with not positive Amount",
			fields: fields{
				FoodId:   "food123",
				SizeId:   "size123",
				Amount:   -5,
				Toppings: &InvoiceDetailToppings{},
			},
			wantErr: true,
		},
		{
			name: "InvoiceDetailCreate is invalid with nil Toppings",
			fields: fields{
				FoodId:   "food123",
				SizeId:   "size123",
				Amount:   2,
				Toppings: nil,
			},
			wantErr: true,
		},
		{
			name: "InvoiceDetailCreate is valid",
			fields: fields{
				FoodId:      "food123",
				FoodName:    "Food Name",
				SizeId:      "size123",
				SizeName:    "Large",
				Amount:      2,
				UnitPrice:   500,
				Toppings:    &InvoiceDetailToppings{},
				Description: "Description",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &InvoiceDetailCreate{
				InvoiceId:   tt.fields.InvoiceId,
				FoodId:      tt.fields.FoodId,
				FoodName:    tt.fields.FoodName,
				SizeId:      tt.fields.SizeId,
				SizeName:    tt.fields.SizeName,
				Toppings:    tt.fields.Toppings,
				Amount:      tt.fields.Amount,
				UnitPrice:   tt.fields.UnitPrice,
				Description: tt.fields.Description,
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
