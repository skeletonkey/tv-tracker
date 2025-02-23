package db

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skeletonkey/lib-core-go/logger"
)

var dbInst *sql.DB

// InitDb sets up the connection to the database and allows for graceful shutdown.
// The wait group is properly incremented.
func InitDb(ctx context.Context, wg *sync.WaitGroup) {
	log := logger.Get()
	log.Info().Msg("Initializing Database")

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		<-ctx.Done()
		if dbInst != nil {
			dbInst.Close()
		}
		log.Info().Msg("Database Closed")
	}(ctx, wg)

	cfg := getConfig()
	var err error
	dbInst, err = sql.Open("sqlite3", cfg.File)
	if err != nil {
		log.Fatal().Err(err).Str("db file name", cfg.File).Msg("error opening db file")
	}
}

func getDb() *sql.DB {
	return dbInst
}
