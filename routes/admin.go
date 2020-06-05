package routes

import (
	"github.com/labstack/echo/v4"
	admin2 "hero/controllers/admin"
	"hero/logger"
)

func InitAdmin(echo *echo.Echo) {
	logger.Info("init admin")
	admin := echo.Group("/admin")
	//record start
	admin.GET("record", admin2.Record)
	//record end
}
