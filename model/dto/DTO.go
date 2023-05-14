package dto

import "database/sql"

type DataIN struct {
	UserID   string
	UserName string
	Tx       *sql.Tx
}

type CustomerIn struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerOut CustomerIn

type UserIn struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
