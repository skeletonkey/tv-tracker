package main

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/skeletonkey/tv-tracker/app/tvdb"
)

func main() {
	e := echo.New()

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