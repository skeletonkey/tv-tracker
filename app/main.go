package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	"github.com/skeletonkey/tv-tracker/app/tvdb"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8081",
			"http://192.168.0.22:8081",
			"http://0.0.0.0:8081",
		},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/search/:query", searchHandler)

	e.Logger.Fatal(e.Start(":8083"))
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