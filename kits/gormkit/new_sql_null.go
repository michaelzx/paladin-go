package gormkit

import (
	"database/sql"
	"time"
)

func NewSqlNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}
