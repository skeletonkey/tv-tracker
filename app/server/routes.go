package server

import (
	echo "github.com/labstack/echo/v4"
)

func setRoutes(e *echo.Echo) {
	e.GET("/search/:query", searchHandler)

	v1Group := e.Group("/api/v1")

	// User
	v1Group.POST("/user", createUser)
	v1Group.GET("/user", getUserId)

}
