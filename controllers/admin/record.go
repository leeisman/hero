package admin

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Record(c echo.Context) error {
	return c.String(http.StatusOK, "/Record")
}
