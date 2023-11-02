package supplierdebtmodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type SupplierDebtType int

const (
	Pay SupplierDebtType = iota
	Debt
)

var allSupplierDebtType = [2]string{"Pay", "Debt"}

func (supplierDebtType *SupplierDebtType) String() string {
	return allSupplierDebtType[*supplierDebtType]
}

func parseStrSupplierDebtType(s string) (SupplierDebtType, error) {
	for i := range allSupplierDebtType {
		if allSupplierDebtType[i] == s {
			return SupplierDebtType(i), nil
		}
	}
	return SupplierDebtType(0), errors.New("invalid debt type string")
}

func (supplierDebtType *SupplierDebtType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStrSupplierDebtType(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*supplierDebtType = v

	return nil
}

func (supplierDebtType *SupplierDebtType) Value() (driver.Value, error) {
	if supplierDebtType == nil {
		return nil, nil
	}

	return supplierDebtType.String(), nil
}

func (supplierDebtType *SupplierDebtType) MarshalJSON() ([]byte, error) {
	if supplierDebtType == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", supplierDebtType.String())), nil
}

func (supplierDebtType *SupplierDebtType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	supplierDebtTypeValue, err := parseStrSupplierDebtType(str)

	if err != nil {
		return err
	}

	*supplierDebtType = supplierDebtTypeValue

	return nil
}
