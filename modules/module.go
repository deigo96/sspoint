package modules

import (
	"transactionType-service/config"
	"transactionType-service/controller"
	"transactionType-service/handler"
	"transactionType-service/helper"
	"transactionType-service/router"
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
