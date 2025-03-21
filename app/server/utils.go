package server

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"

	"github.com/skeletonkey/lib-core-go/logger"
)

var validate = validator.New()

// ValidateAndBind binds and validates a struct from the request context.
func ValidateAndBind(c echo.Context, input interface{}) error {
    log := logger.Get()

    if err := c.Bind(input); err != nil {
        log.Debug().Err(err).Msg("Error Binding")
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    if err := validate.Struct(input); err != nil {
        log.Debug().Err(err).Msg("Error Validating")
        return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
    }

    return nil
}