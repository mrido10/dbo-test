package productDao

import (
	"database/sql"
	"dbo-test/model"
	"dbo-test/util"
	"fmt"
)

func GetProductById(id int64) (prod model.ProductModel, err error) {
	query := fmt.Sprintf(`
		SELECT id, name, price FROM product p 
		WHERE id = $1 `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	err = db.QueryRow(query, id).Scan(&prod.Id.Int64, &prod.Name.String, &prod.Price.Float64)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}