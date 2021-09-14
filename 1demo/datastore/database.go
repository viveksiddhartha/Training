package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBCONN() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "sv_crm:pass#word1@tcp(localhost:3306)/DEMO")
	if err != nil {
		log.Fatal(err)
	}

	return db

}
