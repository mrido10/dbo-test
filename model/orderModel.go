package model

import "database/sql"

type OrderModel struct {
	Id         sql.NullInt64
	CustomerId sql.NullInt64
	ProductId  sql.NullInt64
	TotalOrder sql.NullInt32
	Amount     sql.NullFloat64
	CustomerModel
	ProductModel
}
