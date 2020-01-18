package types

import (
	"database/sql"
	"time"
)

type NullTime sql.NullTime
type NullString sql.NullString
type NullInt32 sql.NullInt32
type NullInt64 sql.NullInt64
type NullFloat64 sql.NullFloat64
type NullBool sql.NullBool

type Time time.Time
