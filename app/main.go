package main

import (
	"context"
	"net/http"
	"sync"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	"github.com/skeletonkey/tv-tracker/app/db"
	"github.com/skeletonkey/tv-tracker/app/tvdb"
)

func main() {
	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)


	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	e := echo.New()
	db.InitDb(ctx, &wg)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8081",
			"http://192.168.0.22:8081",
			"http://0.0.0.0:8081",
		},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/search/:query", searchHandler)

	go func() {

	}
	e.Logger.Fatal(e.Start(":8083"))

	cancel()
	wg.Wait()
}

func searchHandler(c echo.Context) error {
	query := c.Param("query")
	res, err := tvdb.Search(query)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if len(res) == 0 {
		return c.String(http.StatusNoContent, "")
	}
	return c.JSON(http.StatusOK, res)
}
