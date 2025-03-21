package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/tv-tracker/app/db"
	"github.com/skeletonkey/tv-tracker/app/server"
)

func main() {
	log := logger.Get()
	log.Info().Msg("Starting Service")

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGABRT, // Abort - terminate abnormally
		syscall.SIGHUP,  // Hangup - terminal closed or process terminated
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGQUIT, // Ctrl+\
		syscall.SIGTERM, // Terminate - request graceful shutdown
	)

	db.InitDb(ctx, &wg)
	server.Run(ctx, &wg)

	select {
	case sig := <-sigChan:
		log.Info().Int("OS Signal", int(sig.(syscall.Signal))).Msg("OS Signal received")
		cancel()
	case <-ctx.Done():
		log.Info().Msg("context has been cancelled")
	}

	wg.Wait()
	log.Info().Msg("Shutting down service")
}
