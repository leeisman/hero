package routes

import (
	"github.com/labstack/echo/v4"
	admin2 "hero/controllers/admin"
	"hero/pkg/logger"
)

func InitAdmin(echo *echo.Echo) {
	logger.Debug("init admin")
	admin := echo.Group("/admin")
	//record start
	admin.GET("record", admin2.Record)
	//record end
}
