// API server allowing for interactions with the database
package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/skeletonkey/lib-core-go/logger"
)

// Run the web server. The wait group is properly incremented
func Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go start(ctx, wg)
}

func start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	cfg := getConfig()
	log := logger.Get()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.CorsAllow,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	setRoutes(e)

	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("shutting down the server")
		}
	}()

	<-ctx.Done()
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, time.Duration(cfg.ShutdownTimeout)*time.Second)
	defer shutdownCancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Err(err).Msg("server shutdown failed")
	}
	log.Info().Msg("Web Server gracefully shutdown")
}
