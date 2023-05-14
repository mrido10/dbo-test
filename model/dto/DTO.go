package dto

import "database/sql"

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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

type OrderIn struct {
	Id         int64 `json:"id"`
	CustomerId int64 `json:"customer_id"`
	ProductId  int64 `json:"product_id"`
	TotalOrder int32 `json:"total_order"`
}

type OrderOut struct {
	Id           int64   `json:"id"`
	CustomerId   int64   `json:"customer_id"`
	Customername string  `json:"customername"`
	ProductId    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	TotalOrder   int32   `json:"total_order"`
	Amount       float64 `json:"amount"`
}
