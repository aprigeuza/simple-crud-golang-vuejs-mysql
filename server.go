package main

import (
	"crud/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", controllers.HomeIndex())

	e.GET("/contact/list", controllers.ContactGetList())
	e.POST("/contact/add", controllers.ContactAdd())
	e.PUT("/contact/update", controllers.ContactUpdate())
	e.DELETE("/contact/delete/:id", controllers.ContactDelete())

	e.Logger.Fatal(e.Start(":8080"))
}
