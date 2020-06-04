package tracking

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func StartGame(c echo.Context) error {
	return c.String(http.StatusOK, "/StartGame")
}

func Share(c echo.Context) error {
	return c.String(http.StatusOK, "/Share")
}

func DownloadGame(c echo.Context) error {
	return c.String(http.StatusOK, "/DownloadGame")
}
