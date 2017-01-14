package main

import (
	"flag"

	"bitbucket.org/jojahner/page/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func main() {
	flag.Parse()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/", controller.Index)
	e.GET("/stylesheet.css", controller.Stylesheet)

	// start server
	e.Logger.Fatal(e.Start(*addr))
}
