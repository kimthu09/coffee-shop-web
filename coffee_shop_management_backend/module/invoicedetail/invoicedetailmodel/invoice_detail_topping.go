package invoicedetailmodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type InvoiceDetailTopping struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func (data *InvoiceDetailTopping) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var topping InvoiceDetailTopping
	if err := json.Unmarshal(bytes, &topping); err != nil {
		return err
	}

	*data = topping
	return nil
}

func (data *InvoiceDetailTopping) Value() (driver.Value, error) {
	if data == nil {
		return nil, nil
	}

	return json.Marshal(data)
}

type InvoiceDetailToppings []InvoiceDetailTopping

func (data *InvoiceDetailToppings) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB %s", value))
	}

	var toppings []InvoiceDetailTopping
	if err := json.Unmarshal(bytes, &toppings); err != nil {
		return err
	}

	*data = toppings
	return nil
}

func (data *InvoiceDetailToppings) Value() (driver.Value, error) {
	if data == nil {
		return nil, nil
	}

	return json.Marshal(data)
}
