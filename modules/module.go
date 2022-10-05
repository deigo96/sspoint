package modules

import (
	"transactionPoint-service/config"
	"transactionPoint-service/controller"
	"transactionPoint-service/handler"
	"transactionPoint-service/helper"
	"transactionPoint-service/router"
)

func RegisterModules(dbCon *config.DatabaseConnection, c *config.AppConfig) router.Controller {
	refferal := config.RepositoryFactory(dbCon)
	jwtService := handler.NewJWTService()
	referralService := helper.NewTransactionPointService(refferal)
	controller := router.Controller{
		Auth: controller.NewAuthController(referralService, jwtService),
	}

	return controller
}
