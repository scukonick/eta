package structs

import "database/sql"

const (
	StatusNotProcessed int8 = iota
	StatusProcessed
)

type Result struct {
	ID     string        `db:"id"`
	Status int8          `db:"status"`
	ETA    sql.NullInt64 `db:"eta"`
}
