package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var dbInst *sql.DB

func getDb() *sql.DB {
	cfg := getConfig()
	var err error
	dbInst, err = sql.Open("sqlite3", cfg.File)
	if err != nil {
		panic(err)
	}

	return dbInst
}

func closeDb() {
	if dbInst != nil {
		dbInst.Close()
	}
}
