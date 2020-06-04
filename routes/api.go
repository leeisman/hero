package routes

import (
	"github.com/labstack/echo/v4"
	record2 "hero/controllers/api/record"
	"hero/controllers/api/test"
	tracking2 "hero/controllers/api/tracking"
	"hero/logger"
)

func InitApi(echo *echo.Echo) {
	logger.Info("init api")
	api := echo.Group("/api")
	api.GET("/test", test.Test)

	//tracking start
	tracking := api.Group("/tracking")
	tracking.POST("/start_game", tracking2.StartGame)
	tracking.POST("/share", tracking2.Share)
	tracking.POST("/download_game", tracking2.DownloadGame)
	//tracking end

	//record start
	record := api.Group("/record")
	record.POST("game", record2.Record)
	//record end
}
