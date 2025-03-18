package server

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"

	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/tv-tracker/app/db"
	"github.com/skeletonkey/tv-tracker/app/tvdb"
)

func setRoutes(e *echo.Echo) {
	e.GET("/search/:query", searchHandler)

	group := e.Group("/api/v1")

	// User
	group.POST("/user", createUser)
	group.GET("/user", getUser)

}

func createUser(c echo.Context) error {
	log := logger.Get()
	log.Trace().Msg("createUser")
	var user User

	if err := c.Bind(&user); err != nil {
		log.Debug().Err(err).Msg("Error Binding")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Debug().Err(err).Msg("Error Validating")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	userId, err := db.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		log.Debug().Err(err).Msg("Error creating user")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	log.Info().Str("userId", userId).Msg("User created")
	return c.JSON(http.StatusCreated, map[string]string{"user_id": userId})
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

func getUser(c echo.Context) error {
	log := logger.Get()
	log.Trace().Msg("getUser")
	var user User

	if err := c.Bind(&user); err != nil {
		log.Debug().Err(err).Msg("Error Binding")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Debug().Err(err).Msg("Error Validating")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	userId, err := db.GetUser(user.Username, user.Password)
	if err != nil {
		log.Debug().Str("username", user.Username).Err(err).Msg("Error finding user")
		return c.NoContent(http.StatusNotFound)
	}

	log.Info().Str("userId", userId).Msg("User found")
	return c.JSON(http.StatusOK, map[string]string{"user_id": userId})
}
