package types

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
)

// 已测试
func (n *NullFloat64) Scan(value interface{}) error {
	if f64, ok := value.(float64); ok {
		sqlNullFloat64 := sql.NullFloat64{}
		if err := sqlNullFloat64.Scan(f64); err != nil {
			return err
		}
		*n = NullFloat64(sqlNullFloat64)
	}
	return nil
}

func (n NullFloat64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float64, nil
}

// 已测试
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	// return []byte("null"), nil
	if !n.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatFloat(n.Float64, 'f', -1, 64)), nil
}

func (n *NullFloat64) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "null" {
		return nil
	}
	if f64, err := strconv.ParseFloat(dataStr, 64); err != nil {
		return err
	} else {
		n.Float64 = f64
		n.Valid = true
		return nil
	}
}

// func NewNullFloat64(t ...time.Time) NullFloat64 {
// 	if len(t) == 1 {
// 		return NullFloat64(sql.NullFloat64{
// 			Time:  t[0],
// 			Valid: true,
// 		})
// 	} else {
// 		return NullFloat64(sql.NullFloat64{
// 			Time:  time.Time{},
// 			Valid: false,
// 		})
// 	}
// }
