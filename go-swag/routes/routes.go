package routes

import (
	"github.com/joaosoft/swagger/go-swag/controllers"
	"github.com/labstack/echo"
)

var (
	Router = echo.New()
)

func init() {

	Router.GET("/v1/persons/:id_person", controllers.GetPersonByID)
	Router.GET("/v1/persons/:id_person/addresses/:id_address", controllers.GetPersonAddressByID)
	Router.GET("/v1/errors", controllers.GetErrorByID)

	Router.Static("swaggerui", "./spec")

	Router.Logger.Fatal(Router.Start(":8081"))
}
