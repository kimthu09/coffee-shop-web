package productmodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Size struct {
	Id       string  `json:"id" gorm:"column:id;"`
	Name     string  `json:"name" gorm:"column:name;"`
	Cost     float64 `json:"cost" gorm:"column:cost;"`
	Price    float64 `json:"price" gorm:"column:price;"`
	RecipeId string  `json:"recipeId" gorm:"column:recipeId;"`
}

func (size *Size) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var productSize Size
	if err := json.Unmarshal(bytes, &productSize); err != nil {
		return err
	}

	*size = productSize
	return nil
}

func (size *Size) Value() (driver.Value, error) {
	if size == nil {
		return nil, nil
	}

	return json.Marshal(size)
}
