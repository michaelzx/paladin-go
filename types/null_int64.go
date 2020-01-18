package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

func (n *NullInt64) Scan(value interface{}) error {
	if value == nil {
		n.Int64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	switch vt := value.(type) {
	case int64:
		n.Int64 = vt
	default:
		// TODO 待验证
		fmt.Printf("%#v %T", vt, vt)
	}
	return nil
}
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(n.Int64, 10)), nil
}

func (n *NullInt64) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "null" {
		return nil
	}
	if i64, err := strconv.ParseInt(dataStr, 10, 64); err != nil {
		return err
	} else {
		n.Int64 = i64
		n.Valid = true
		return nil
	}
}
