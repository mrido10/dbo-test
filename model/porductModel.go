package model

import "database/sql"

type ProductModel struct {
	Id    sql.NullInt64
	Name  sql.NullString
	Price sql.NullFloat64
}
