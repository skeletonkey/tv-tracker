package server

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"

	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/tv-tracker/app/db"
)

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

func getUserId(c echo.Context) error {
	log := logger.Get()
	log.Trace().Msg("getUserId")
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

	userId, err := db.GetUserId(user.Username, user.Password)
	if err != nil {
		log.Debug().Str("username", user.Username).Err(err).Msg("Error finding user")
		return c.NoContent(http.StatusNotFound)
	}

	log.Info().Str("userId", userId).Msg("User found")
	return c.JSON(http.StatusOK, map[string]string{"user_id": userId})
}
