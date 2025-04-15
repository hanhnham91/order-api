package util

import (
	"errors"
	"net/http"

	pkgerror "github.com/hanhnham91/pkg/error"
	"github.com/labstack/echo/v4"
)

type response struct{}

var Response response

func (response) Success(c echo.Context, data interface{}) error {
	if data != nil {
		return c.JSON(http.StatusOK, data)
	}

	return c.NoContent(http.StatusOK)
}

func (response) Error(c echo.Context, err error) error {
	var appError pkgerror.MyError

	if !errors.As(err, &appError) {
		return c.JSON(appError.HTTPCode, map[string]interface{}{
			"code":    "500-***",
			"message": err.Error(),
		})
	}

	return c.JSON(appError.HTTPCode, map[string]interface{}{
		"code":    appError.ErrorCode,
		"message": appError.Message,
	})
}
