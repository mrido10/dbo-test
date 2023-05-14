package model

import "database/sql"

type CustomerModel struct {
	Id      sql.NullInt64
	Name    sql.NullString
	Email   sql.NullString
	Created sql.NullTime
	Updated sql.NullTime
	Deleted sql.NullBool
}
