package controller

import (
	"github.com/jojahner/website/templates"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	p := &templates.Index{}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	templates.WritePageTemplate(c.Response(), p)
	return nil
}
