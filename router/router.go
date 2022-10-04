package router

import (
	"referralUser-service/controller"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *controller.AuthController
}

func Routes(e *echo.Echo, controller *Controller) {
	e.GET("/referralTree", controller.Auth.GetReferralUser)
	e.POST("/referralTree", controller.Auth.RegisterReferral)
	e.PUT("/referralTree", controller.Auth.UpdateReferral)
}
