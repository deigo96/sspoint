package modules

import (
	"pointHistory-service/config"
	"pointHistory-service/controller"
	"pointHistory-service/handler"
	"pointHistory-service/helper"
	"pointHistory-service/router"
)

func RegisterModules(dbCon *config.DatabaseConnection, c *config.AppConfig) router.Controller {
	refferal := config.RepositoryFactory(dbCon)
	jwtService := handler.NewJWTService()
	pointService := helper.NewRewardService(refferal)
	controller := router.Controller{
		Auth: controller.NewAuthController(pointService, jwtService),
	}

	return controller
}
