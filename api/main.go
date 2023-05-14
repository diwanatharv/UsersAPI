package main

import (
	"github.com/labstack/echo/v4"

	"awesomeProject3/api/controller"
)

func main() {
	e := echo.New()
	controller.Createroutes(e)
	e.Logger.Fatal(e.Start(":7000"))
}
