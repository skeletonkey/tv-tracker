package server

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"

	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/tv-tracker/app/tvdb"
)

func searchHandler(c echo.Context) error {
	log := logger.Get()
	log.Trace().Msg("searchHandler")

	query := c.Param("query")
	no_cache, _ := strconv.ParseBool(c.QueryParam("no_cache"))

	res, err := tvdb.Search(query, no_cache)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if len(res) == 0 {
		return c.String(http.StatusNoContent, "")
	}
	return c.JSON(http.StatusOK, res)
}
