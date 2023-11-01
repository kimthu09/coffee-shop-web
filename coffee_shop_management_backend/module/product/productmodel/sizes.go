package productmodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Sizes []Size

func (sizes *Sizes) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB %s", value))
	}

	var productSizes []Size
	if err := json.Unmarshal(bytes, &productSizes); err != nil {
		return err
	}

	*sizes = productSizes
	return nil
}

func (sizes *Sizes) Value() (driver.Value, error) {
	if sizes == nil {
		return nil, nil
	}

	return json.Marshal(sizes)
}
