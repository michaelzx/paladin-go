package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"
)

func (n *NullTime) Scan(value interface{}) error {
	if stdTime, ok := value.(time.Time); ok {
		sqlNullTime := sql.NullTime{}
		err := sqlNullTime.Scan(stdTime)
		if err != nil {
			return err
		}
		*n = NullTime(sqlNullTime)
	}
	return nil
}

func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
func (n NullTime) String() (string, error) {
	if n.Valid {
		return n.Time.Format("2006-01-02 15:04:05"), nil
	} else {
		return "n.Time", errors.New("invalid")
	}
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	// return []byte("null"), nil
	if !n.Valid {
		return []byte("null"), nil
	}
	str := n.Time.Format(`"2006-01-02 15:04:05"`)
	return []byte(str), nil
}
func (n *NullTime) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "null" {
		return nil
	}
	var err error
	stdTime, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), TimeZone)
	if err != nil {
		return nil
	}
	n.Time = stdTime
	n.Valid = true
	return nil
}
func NewNullTime(t ...time.Time) NullTime {
	if len(t) == 1 {
		return NullTime(sql.NullTime{
			Time:  t[0],
			Valid: true,
		})
	} else {
		return NullTime(sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		})
	}
}
