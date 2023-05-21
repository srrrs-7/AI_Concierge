package entity

import "database/sql"

type Author struct {
	ID   int64
	Name string
	Bio  sql.NullString
}
