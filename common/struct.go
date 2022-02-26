package common

import "database/sql"

var Db *sql.DB

type Record struct {
	Id       string
	Name     string
	Position string
}
