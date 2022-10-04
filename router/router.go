package router

import (
	"reward-list-service/controller"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *controller.AuthController
}

func Routes(e *echo.Echo, controller *Controller) {
	e.GET("/rewardList", controller.Auth.GetReward)
	e.POST("/rewardList", controller.Auth.StoreReward)
	e.PUT("/rewardList", controller.Auth.UpdateReward)
}
