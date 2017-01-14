package controller

import (
	"bitbucket.org/jojahner/page/templates"
	"github.com/labstack/echo"
)

func Stylesheet(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/css")
	templates.WriteStylesheet(c.Response())
	return nil
}
