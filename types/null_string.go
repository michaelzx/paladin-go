package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func (n *NullString) Scan(value interface{}) error {
	if value == nil {
		n.String, n.Valid = "", false
		return nil
	}
	n.Valid = true
	switch vt := value.(type) {
	case string:
		n.String = vt
	case []byte:
		n.String = string(vt)
	default:
		// TODO 待验证
		fmt.Printf("NullString %#v %T", vt, vt)
	}
	return nil
}
func (n NullString) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}
func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	stringJsonBytes, err := json.Marshal(n.String)
	if err != nil {
		return nil, err
	}
	return stringJsonBytes, nil
}

func (n *NullString) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "null" {
		return nil
	}
	n.String = dataStr
	n.Valid = true
	return nil
}
