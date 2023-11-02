package cancelnotedetailmodel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type CancelReason int

const (
	Damaged CancelReason = iota
	OutOfDate
)

var allCancelReason = [2]string{"Damaged", "OutOfDate"}

func (cancelReason *CancelReason) String() string {
	return allCancelReason[*cancelReason]
}

func parseStrCancelReason(s string) (CancelReason, error) {
	for i := range allCancelReason {
		if allCancelReason[i] == s {
			return CancelReason(i), nil
		}
	}
	return CancelReason(0), errors.New("invalid cancel reason string")
}

func (cancelReason *CancelReason) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStrCancelReason(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*cancelReason = v

	return nil
}

func (cancelReason *CancelReason) Value() (driver.Value, error) {
	if cancelReason == nil {
		return nil, nil
	}

	return cancelReason.String(), nil
}

func (cancelReason *CancelReason) MarshalJSON() ([]byte, error) {
	if cancelReason == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", cancelReason.String())), nil
}

func (cancelReason *CancelReason) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	cancelReasonValue, err := parseStrCancelReason(str)

	if err != nil {
		return err
	}

	*cancelReason = cancelReasonValue

	return nil
}
