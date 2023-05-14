package main

import (
	"dbo-test/app"
	"dbo-test/config"
	"dbo-test/util"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"log"
)

func init() {
	var err error
	config.Configure, err = config.GetConfig("config/config.yml")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	migrateSQL()
	app.Start()
}

func migrateSQL() {
	db, err := util.ConnectPostgres()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	n, err := migrate.Exec(db, "postgres", &migrate.FileMigrationSource{Dir: "sql/"}, migrate.Up)
	if err != nil {
		log.Fatal("Error occcured:", err)
		return
	}

	log.Println(fmt.Sprintf("Applied %d migrations!\n", n))
}
