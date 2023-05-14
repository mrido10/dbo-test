package userDao

import (
	"database/sql"
	"dbo-test/model"
	"dbo-test/util"
	"fmt"
)

func GetUser(userName, password string) (user model.UserModel, err error) {
	query := fmt.Sprintf(`
		SELECT id, user_name
		FROM "user" u 
		WHERE user_name = $1
		AND password = $2 `)
	db, err := util.ConnectPostgres()
	if err != nil {
		return
	}
	defer db.Close()
	params := []interface{}{userName, password}

	err = db.QueryRow(query, params...).Scan(&user.Id.Int64, &user.UseName.String)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	err = nil
	return
}
