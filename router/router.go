package router

import (
	"referralUser-service/controller"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *controller.AuthController
}

func Routes(e *echo.Echo, controller *Controller) {
	e.POST("/trxType", controller.Auth.StoreTransactionType)
	e.PUT("/trxType", controller.Auth.UpdateTransactionType)
}
