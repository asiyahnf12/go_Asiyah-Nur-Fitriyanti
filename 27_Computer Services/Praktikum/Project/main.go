package main

import (
	"mini_project/config"
	"mini_project/middleware"
	"mini_project/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	middleware.Logmiddleware(e)

	db := config.Init()

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
