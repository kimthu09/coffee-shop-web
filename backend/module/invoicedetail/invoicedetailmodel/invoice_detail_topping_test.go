package invoicedetailmodel

import (
	"database/sql/driver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceDetailTopping_Scan(t *testing.T) {
	type fields struct {
		Id    string
		Name  string
		Price int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantedData InvoiceDetailTopping
		wantErr    bool
	}{
		{
			name: "Scan successfully unmarshal JSONB",
			args: args{
				value: []byte(
					`{"id": "topping123","name": "Topping Name","price": 10000}`,
				),
			},
			wantedData: InvoiceDetailTopping{
				Id:    "topping123",
				Name:  "Topping Name",
				Price: 10000,
			},
			wantErr: false,
		},
		{
			name:       "Scan fails with non-JSONB value",
			args:       args{value: []byte("invalid-value")},
			wantedData: InvoiceDetailTopping{},
			wantErr:    true,
		},
		{
			name:       "Scan fails with invalid JSONB format",
			args:       args{value: []byte(`{"id": "topping123", "price": 5,}`)},
			wantedData: InvoiceDetailTopping{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := &InvoiceDetailTopping{
				Id:    tt.fields.Id,
				Name:  tt.fields.Name,
				Price: tt.fields.Price,
			}
			err := data.Scan(tt.args.value)

			if tt.wantErr {
				assert.NotNil(t, err, "Scan() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Scan() = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(
					t, *data, tt.wantedData,
					"data = ?, want = ?",
					*data, tt.wantedData.Id)
			}
		})
	}
}

func TestInvoiceDetailTopping_Value(t *testing.T) {
	tests := []struct {
		name    string
		topping *InvoiceDetailTopping
		want    driver.Value
		wantErr bool
	}{
		{
			name: "Value successfully marshals JSONB",
			topping: &InvoiceDetailTopping{
				Id:    "topping123",
				Name:  "Topping Name",
				Price: 10000,
			},
			want: []byte(
				`{"id":"topping123","name":"Topping Name","price":10000}`,
			),
			wantErr: false,
		},
		{
			name:    "Value successfully marshals nil",
			topping: nil,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.topping.Value()

			if tt.wantErr {
				assert.NotNil(t, err, "Value() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Value() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvoiceDetailToppings_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		data       InvoiceDetailToppings
		args       args
		wantedData InvoiceDetailToppings
		wantErr    bool
	}{
		{
			name: "Scan successfully unmarshal JSONB",
			data: InvoiceDetailToppings{},
			args: args{
				value: []byte(`[{"id":"topping1","name":"Topping 1","price":100},{"id":"topping2","name":"Topping 2","price":200}]`),
			},
			wantedData: InvoiceDetailToppings{
				{
					Id:    "topping1",
					Name:  "Topping 1",
					Price: 100,
				},
				{
					Id:    "topping2",
					Name:  "Topping 2",
					Price: 200,
				},
			},
			wantErr: false,
		},
		{
			name: "Scan failed with non-JSONB",
			data: InvoiceDetailToppings{},
			args: args{
				value: "invalid-json",
			},
			wantedData: InvoiceDetailToppings{},
			wantErr:    true,
		},
		{
			name: "Scan failed with invalid JSONB format",
			data: InvoiceDetailToppings{},
			args: args{
				value: []byte(`"invalid-json": "123"`),
			},
			wantedData: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.data.Scan(tt.args.value)

			if tt.wantErr {
				assert.NotNil(t, err, "Scan() = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Scan() = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(
					t, tt.data, tt.wantedData,
					"data = ?, want = ?",
					tt.data, tt.wantedData)
			}
		})
	}
}

func TestInvoiceDetailToppings_Value(t *testing.T) {
	tests := []struct {
		name    string
		data    *InvoiceDetailToppings
		want    driver.Value
		wantErr bool
	}{
		{
			name: "Value successfully marshals JSONB",
			data: &InvoiceDetailToppings{
				{
					Id:    "topping1",
					Name:  "Topping 1",
					Price: 100,
				},
				{
					Id:    "topping2",
					Name:  "Topping 2",
					Price: 200,
				},
			},
			want:    []byte(`[{"id":"topping1","name":"Topping 1","price":100},{"id":"topping2","name":"Topping 2","price":200}]`),
			wantErr: false,
		},
		{
			name:    "Value successfully handles nil",
			data:    nil,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.data.Value()

			if tt.wantErr {
				assert.NotNil(t, err, "Value() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "Value() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, got, tt.want, "Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
