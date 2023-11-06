package enum

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type MeasureType int

const (
	Weight MeasureType = iota
	Volume
	Unit
)

var allMeasureType = [3]string{"Weight", "Volume", "Unit"}

func (measureType *MeasureType) String() string {
	return allMeasureType[*measureType]
}

func parseStrMeasureType(s string) (MeasureType, error) {
	for i := range allMeasureType {
		if allMeasureType[i] == s {
			return MeasureType(i), nil
		}
	}
	return MeasureType(0), errors.New("invalid measure type string")
}

func (measureType *MeasureType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStrMeasureType(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*measureType = v

	return nil
}

func (measureType *MeasureType) Value() (driver.Value, error) {
	if measureType == nil {
		return nil, nil
	}

	return measureType.String(), nil
}

func (measureType *MeasureType) MarshalJSON() ([]byte, error) {
	if measureType == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", measureType.String())), nil
}

func (measureType *MeasureType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	measureTypeValue, err := parseStrMeasureType(str)

	if err != nil {
		return err
	}

	*measureType = measureTypeValue

	return nil
}
