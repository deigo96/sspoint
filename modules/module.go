package modules

import (
	"reward-list-service/config"
	"reward-list-service/controller"
	"reward-list-service/handler"
	"reward-list-service/helper"
	"reward-list-service/router"
)

func RegisterModules(dbCon *config.DatabaseConnection, c *config.AppConfig) router.Controller {
	reward := config.RepositoryFactory(dbCon)
	jwtService := handler.NewJWTService()
	rewardService := helper.NewRewardService(reward)
	controller := router.Controller{
		Auth: controller.NewAuthController(rewardService, jwtService),
	}

	return controller
}
