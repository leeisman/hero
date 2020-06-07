package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hero/pkg/logger"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)
		logger.Print(c.Request().URL, "reqBody", requestID, string(reqBody))
		logger.Print(c.Request().URL, "resBody", requestID, string(resBody))
	}))
	InitApi(e)
	InitAdmin(e)
	e.Logger.Fatal(e.Start(":9000"))
}
