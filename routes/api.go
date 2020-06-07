package routes

import (
	"github.com/labstack/echo/v4"
	"hero/controllers/api/active/hero"
	"hero/pkg/logger"
)

func InitApi(echo *echo.Echo) {
	logger.Debug("init api")
	api := echo.Group("/api")
	initActiveHero(api)
}

func initActiveHero(api *echo.Group) {
	activeHero := api.Group("/active/hero")
	activeHero.POST("/play", hero.Play)
	activeHero.POST("/record", hero.Record)
	activeHero.POST("/tracking", hero.Tracking)
}
