package db

import (
	"context"
	"database/sql"
	"os"
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
	dbFile := "../" + cfg.File
	if _, err := os.Stat(dbFile); err != nil {
		if err != nil {
			log.Fatal().Err(err).Str("dbFile", dbFile).Msg("Issues with DB File")
		}
	}

	var err error
	dbInst, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal().Err(err).Str("db file name", cfg.File).Msg("error opening db file")
	}

	if err = dbInst.Ping(); err != nil {
		log.Fatal().Err(err).Msg("unable to ping db")
	}
}

func getDb() *sql.DB {
	return dbInst
}
