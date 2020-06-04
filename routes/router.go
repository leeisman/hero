package routes

import (
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	InitApi(e)
	InitAdmin(e)
	e.Logger.Fatal(e.Start(":9000"))
}
