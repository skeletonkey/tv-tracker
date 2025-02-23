package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var dbInst *sql.DB

func InitDb(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("Initializing Database")
	go func (ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		<-ctx.Done()
		closeDb()
		fmt.Println("Database Closed")
	}(ctx, wg)
	_ = getDb()
}

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
