package main

import (
	"code_competence/config"
	"code_competence/middleware"
	"code_competence/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	middleware.Logmiddleware(e)

	db := config.Init()

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
