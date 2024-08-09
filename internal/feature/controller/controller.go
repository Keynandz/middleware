package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	Get(c echo.Context) error
}
