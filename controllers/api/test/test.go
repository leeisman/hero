package test

import (
	"github.com/labstack/echo/v4"
	"hero/logger"
	"net/http"
)

func Test(c echo.Context) error {
	logger.Info("test")
	return c.String(http.StatusOK, "/users")
}
