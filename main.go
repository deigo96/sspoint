package main

import (
	"net/http"
	"os"
	"os/signal"
	"reward-list-service/config"
	"reward-list-service/modules"
	"reward-list-service/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	con := config.GetConfig()
	dbCon := config.NewDatabaseConnection(con)
	controllers := modules.RegisterModules(dbCon, con)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	router.Routes(e, &controllers)

	server := config.GetServer()

	e.Logger.Fatal(e.Start("127.0.0.1" + server.Port))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
