package server

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	validator "github.com/go-playground/validator/v10"

	"github.com/skeletonkey/tv-tracker/app/tvdb"
)

func setRoutes(e *echo.Echo) {
	e.GET("/search/:query", searchHandler)

	group := e.Group("/api/v1")

	// Users
	group.POST("/user", createUser)


}

func createUser(c echo.Context) error {
	var user User
	validate := validator.New()

	
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
