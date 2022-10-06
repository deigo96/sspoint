package router

import (
	"pointHistory-service/controller"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *controller.AuthController
}

func Routes(e *echo.Echo, controller *Controller) {
	e.GET("/pointHistory", controller.Auth.GetPointHistoryService)
	e.POST("/pointHistory", controller.Auth.StorePointHistoryService)
}
