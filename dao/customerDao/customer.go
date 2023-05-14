package customerDao

import (
	"database/sql"
	"dbo-test/model"
	"dbo-test/util"
	"fmt"
	"strings"
)

func GetCustomerByUnique(name, email string, isCheckDeleted bool) (customer model.CustomerModel, err error) {
	query := fmt.Sprintf(`
		SELECT id, name, email 
		FROM customer 
		WHERE name = $1
			AND email = $2 `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	params := []interface{}{name, email}
	if isCheckDeleted {
		query += " AND deleted = $3 "
		params = append(params, !isCheckDeleted)
	}

	err = db.QueryRow(query, params...).Scan(&customer.Id.Int64, &customer.Name.String, &customer.Email.String)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

func GetCustomerById(id int64, isCheckDeleted bool) (customer model.CustomerModel, err error) {
	query := fmt.Sprintf(`
		SELECT id, name, email 
		FROM customer 
		WHERE id = $1 `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	params := []interface{}{id}
	if isCheckDeleted {
		query += " AND deleted = $2 "
		params = append(params, !isCheckDeleted)
	}

	err = db.QueryRow(query, params...).Scan(&customer.Id.Int64, &customer.Name.String, &customer.Email.String)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}

func UpsertCustomer(tx *sql.Tx, data model.CustomerModel) (err error) {
	query := fmt.Sprintf(`
		INSERT INTO customer (name, email)
		VALUES($1, $2)
		ON CONFLICT(name, email)
		DO UPDATE 
			SET name = EXCLUDED.name,
			email = EXCLUDED.email,
			created = now(),
			updated = now(),
			deleted = FALSE `)
	param := []interface{}{data.Name.String, data.Email.String}
	_, err = tx.Exec(query, param...)
	return
}

func UpdateCustomer(tx *sql.Tx, data model.CustomerModel) (err error) {
	query := fmt.Sprintf(`
		UPDATE customer 
		SET name = $1,
			email = $2,
			updated = now(),
			deleted = FALSE 
		WHERE id = $3 `)
	param := []interface{}{data.Name.String, data.Email.String, data.Id.Int64}
	_, err = tx.Exec(query, param...)
	return
}

func GetListCustomer(page int, search string) ([]model.CustomerModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			id, name, email, created, updated, deleted 
		FROM customer c 
		WHERE deleted = FALSE `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if strings.TrimSpace(search) != "" {
		query += " AND lower(name) LIKE lower('%" + search + "%') "
	}

	query += " ORDER BY updated DESC LIMIT 10 OFFSET $1 "

	rows, err := db.Query(query, (page-1)*10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.CustomerModel
	for rows.Next() {
		var temp model.CustomerModel
		err = rows.Scan(
			&temp.Id.Int64, &temp.Name.String, &temp.Email.String, &temp.Created.Time, &temp.Updated.Time, &temp.Deleted,
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

func DeleteCustomer(tx *sql.Tx, id int64) (err error) {
	query := fmt.Sprintf(`
		UPDATE customer 
		SET deleted = TRUE,
			updated = now()
		WHERE id = $1 `)
	_, err = tx.Exec(query, id)
	return
}
