package router

import (
	"transactionPoint-service/controller"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *controller.AuthController
}

func Routes(e *echo.Echo, controller *Controller) {
	e.POST("/trxPoint", controller.Auth.StoreTransactionPoint)
	e.PUT("/trxPoint", controller.Auth.UpdateTransactionPoint)
}
