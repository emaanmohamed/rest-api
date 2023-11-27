package main

import (
	"github.com/emaanmohamed/rest-api/models"
	"github.com/emaanmohamed/rest-api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	models.ConnectToDB()

}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	routes.SetUpRoutes(api)

	e.Logger.Fatal(e.Start(":8099"))
}
