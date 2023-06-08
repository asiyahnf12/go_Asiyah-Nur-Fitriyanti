package main

import (
	"go-latihan/config"
	"go-latihan/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	middlewares.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":8000"))
}
