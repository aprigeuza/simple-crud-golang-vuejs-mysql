package controllers

import (
	"github.com/labstack/echo/v4"
)

func HomeIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.File("public/index.html")
	}
}
