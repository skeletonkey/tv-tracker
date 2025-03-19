package server

import (
	echo "github.com/labstack/echo/v4"
)

func setRoutes(e *echo.Echo) {
	e.GET("/search/:query", searchHandler)

	group := e.Group("/api/v1")

	// User
	group.POST("/user", createUser)
	group.GET("/user", getUser)

}
