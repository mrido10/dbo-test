package model

import "database/sql"

type UserModel struct {
	Id       sql.NullInt64
	UseName  sql.NullString
	Password sql.NullString
}
