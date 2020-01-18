package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

func (n *NullInt32) Scan(value interface{}) error {
	if value == nil {
		n.Int32, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	switch vt := value.(type) {
	case int64:
		n.Int32 = int32(vt)
	default:
		// TODO 待验证
		fmt.Printf("NullInt32 %#v %T", vt, vt)
	}
	return nil
}
func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int32), nil
}
func (n NullInt32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return []byte(formatInt32(n.Int32)), nil
}

func (n *NullInt32) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "null" {
		return nil
	}
	if i32, err := strconv.ParseInt(dataStr, 10, 32); err != nil {
		return err
	} else {
		n.Int32 = int32(i32)
		n.Valid = true
		return nil
	}
}
func formatInt32(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
