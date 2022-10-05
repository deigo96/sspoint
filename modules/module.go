package modules

import (
	"referralUser-service/config"
	"referralUser-service/controller"
	"referralUser-service/handler"
	"referralUser-service/helper"
	"referralUser-service/router"
)

func RegisterModules(dbCon *config.DatabaseConnection, c *config.AppConfig) router.Controller {
	refferal := config.RepositoryFactory(dbCon)
	jwtService := handler.NewJWTService()
	referralService := helper.NewTransactionTypeService(refferal)
	controller := router.Controller{
		Auth: controller.NewAuthController(referralService, jwtService),
	}

	return controller
}
