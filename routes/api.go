package routes

import (
	"github.com/labstack/echo/v4"
	"hero/controllers/api/active/hero"
	report "hero/controllers/api/report/hero"
	"hero/pkg/logger"
)

func InitApi(echo *echo.Echo) {
	logger.Debug("init api")
	api := echo.Group("/api")
	initActiveHero(api)
	initReportHero(api)
}

func initActiveHero(api *echo.Group) {
	activeHero := api.Group("/active/hero")
	activeHero.POST("/play", hero.Play)
	activeHero.POST("/record", hero.Record)
	activeHero.POST("/prize", hero.Prize)
	activeHero.POST("/tracking", hero.Tracking)

}

func initReportHero(api *echo.Group) {
	reportHero := api.Group("/report/hero")
	reportHero.GET("/user_count", report.UserCount)
	reportHero.GET("/score_count", report.ScoreCount)
	reportHero.GET("/rank", report.Rank)
}
