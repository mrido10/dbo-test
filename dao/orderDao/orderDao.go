package orderDao

import (
	"database/sql"
	"dbo-test/model"
	"dbo-test/util"
	"fmt"
	"strings"
)

func InsertOrder(tx *sql.Tx, data model.OrderModel) (err error) {
	query := fmt.Sprintf(`
		INSERT INTO "order"(customer_id, product_id, total_order, amount)
		VALUES($1, $2, $3, $4) `)
	param := []interface{}{data.CustomerId.Int64, data.ProductId.Int64, data.TotalOrder.Int32, data.Amount.Float64}
	_, err = tx.Exec(query, param...)
	return
}

func UpdateOrder(tx *sql.Tx, data model.OrderModel) (err error) {
	query := fmt.Sprintf(`
		UPDATE "order"
		SET customer_id = $1, 
			product_id = $2,
			total_order = $3, 
			amount = $4,
			updated = now(),
			deleted = FALSE 
		WHERE id = $5`)
	param := []interface{}{data.CustomerId.Int64, data.ProductId.Int64, data.TotalOrder.Int32, data.Amount.Float64,
		data.Id.Int64}
	_, err = tx.Exec(query, param...)
	return
}

func GetListOrder(page int, searchCustomer, searchProduct string) ([]model.OrderModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			o.id, o.customer_id, c.name, o.product_id, p.name, o.total_order, o.amount 
		FROM "order" o
		INNER JOIN customer c 
			ON c.id = o.customer_id 
		INNER JOIN product p 
			ON p.id = o.id 
		WHERE o.deleted = FALSE 
			AND c.deleted = FALSE 
			AND p.deleted = FALSE `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if strings.TrimSpace(searchCustomer) != "" {
		query += " AND lower(c.name) LIKE lower('%"+searchCustomer+"%') "
	}
	if strings.TrimSpace(searchProduct) != "" {
		query += " AND lower(p.name) LIKE lower('%"+searchProduct+"%') "
	}

		query += " ORDER BY o.updated DESC LIMIT 10 OFFSET $1 "

	rows, err := db.Query(query, (page-1)*10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.OrderModel
	for rows.Next() {
		var temp model.OrderModel
		err = rows.Scan(
			&temp.Id.Int64, &temp.CustomerId.Int64, &temp.CustomerModel.Name.String, &temp.ProductId.Int64,
			&temp.ProductModel.Name.String, &temp.TotalOrder.Int32, &temp.Amount.Float64,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, temp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, err
}

func CountOrderById(id int64) (count int, err error) {
	query := fmt.Sprintf(`
		SELECT 
			count(id)
		FROM "order" 
		WHERE id = $1 `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	err = db.QueryRow(query, id).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

func GetOrderById(id int64) (od model.OrderModel, err error) {
	query := fmt.Sprintf(`
		SELECT 
			o.id, o.customer_id, c.name, o.product_id, p.name, o.total_order, o.amount 
		FROM "order" o
		INNER JOIN customer c 
			ON c.id = o.customer_id 
		INNER JOIN product p 
			ON p.id = o.id 
		WHERE o.id = $1 
			AND o.deleted = FALSE 
			AND c.deleted = FALSE 
			AND p.deleted = FALSE `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	err = db.QueryRow(query, id).Scan(&od.Id.Int64, &od.CustomerId.Int64, &od.CustomerModel.Name.String, &od.ProductId.Int64,
		&od.ProductModel.Name.String, &od.TotalOrder.Int32, &od.Amount.Float64,)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

func DeleteOrder(tx *sql.Tx, id int64) (err error) {
	query := fmt.Sprintf(`
		UPDATE "order" 
		SET deleted = TRUE,
			updated = now()
		WHERE id = $1 `)
	_, err = tx.Exec(query, id)
	return
}