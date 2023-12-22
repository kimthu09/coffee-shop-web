package invoicemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceCreate_TableName(t *testing.T) {
	type fields struct {
		Id                  string
		CustomerId          *string
		Customer            SimpleCustomer
		TotalPrice          int
		IsUsePoint          bool
		AmountReceived      int
		AmountPriceUsePoint int
		CreatedBy           string
		InvoiceDetails      []invoicedetailmodel.InvoiceDetailCreate
		MapIngredient       map[string]int
		ShopName            string
		ShopPhone           string
		ShopAddress         string
		ShopPassWifi        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get TableName of InvoiceCreate successfully",
			fields: fields{
				Id:                  "",
				CustomerId:          nil,
				Customer:            SimpleCustomer{},
				TotalPrice:          0,
				IsUsePoint:          false,
				AmountReceived:      0,
				AmountPriceUsePoint: 0,
				CreatedBy:           "",
				InvoiceDetails:      nil,
				MapIngredient:       nil,
				ShopName:            "",
				ShopPhone:           "",
				ShopAddress:         "",
				ShopPassWifi:        "",
			},
			want: common.TableInvoice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoice := &InvoiceCreate{
				Id:                  tt.fields.Id,
				CustomerId:          tt.fields.CustomerId,
				Customer:            tt.fields.Customer,
				TotalPrice:          tt.fields.TotalPrice,
				IsUsePoint:          tt.fields.IsUsePoint,
				AmountReceived:      tt.fields.AmountReceived,
				AmountPriceUsePoint: tt.fields.AmountPriceUsePoint,
				CreatedBy:           tt.fields.CreatedBy,
				InvoiceDetails:      tt.fields.InvoiceDetails,
				MapIngredient:       tt.fields.MapIngredient,
				ShopName:            tt.fields.ShopName,
				ShopPhone:           tt.fields.ShopPhone,
				ShopAddress:         tt.fields.ShopAddress,
				ShopPassWifi:        tt.fields.ShopPassWifi,
			}
			got := invoice.TableName()

			assert.Equal(t, tt.want, got, "TableName() = %v, wantedData %v", got, tt.want)
		})
	}
}

func TestInvoiceCreate_Validate(t *testing.T) {
	type fields struct {
		Id                  string
		CustomerId          *string
		Customer            SimpleCustomer
		TotalPrice          int
		IsUsePoint          bool
		AmountReceived      int
		AmountPriceUsePoint int
		CreatedBy           string
		InvoiceDetails      []invoicedetailmodel.InvoiceDetailCreate
		MapIngredient       map[string]int
		ShopName            string
		ShopPhone           string
		ShopAddress         string
		ShopPassWifi        string
	}

	invalidCustomerId := "12345678901234567890"
	emptyDetails := make([]invoicedetailmodel.InvoiceDetailCreate, 0)
	invalidDetail := []invoicedetailmodel.InvoiceDetailCreate{
		{
			FoodId: "",
		},
	}
	validDetail := []invoicedetailmodel.InvoiceDetailCreate{
		{
			FoodId: "Food001",
			SizeId: "Size001",
			Toppings: &invoicedetailmodel.InvoiceDetailToppings{
				{
					Id: "Topping001",
				},
			},
			Amount: 1,
		},
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "InvoiceCreate is invalid with invalid customer id",
			fields: fields{
				CustomerId: &invalidCustomerId,
			},
			wantErr: true,
		},
		{
			name: "InvoiceCreate is invalid with details not exist",
			fields: fields{
				CustomerId:     nil,
				InvoiceDetails: nil,
			},
			wantErr: true,
		},
		{
			name: "InvoiceCreate is invalid with empty details",
			fields: fields{
				CustomerId:     nil,
				InvoiceDetails: emptyDetails,
			},
			wantErr: true,
		},
		{
			name: "InvoiceCreate is invalid with invalid details",
			fields: fields{
				CustomerId:     nil,
				InvoiceDetails: invalidDetail,
			},
			wantErr: true,
		},
		{
			name: "InvoiceCreate is valid",
			fields: fields{
				CustomerId:     nil,
				InvoiceDetails: validDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &InvoiceCreate{
				Id:                  tt.fields.Id,
				CustomerId:          tt.fields.CustomerId,
				Customer:            tt.fields.Customer,
				TotalPrice:          tt.fields.TotalPrice,
				IsUsePoint:          tt.fields.IsUsePoint,
				AmountReceived:      tt.fields.AmountReceived,
				AmountPriceUsePoint: tt.fields.AmountPriceUsePoint,
				CreatedBy:           tt.fields.CreatedBy,
				InvoiceDetails:      tt.fields.InvoiceDetails,
				MapIngredient:       tt.fields.MapIngredient,
				ShopName:            tt.fields.ShopName,
				ShopPhone:           tt.fields.ShopPhone,
				ShopAddress:         tt.fields.ShopAddress,
				ShopPassWifi:        tt.fields.ShopPassWifi,
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
