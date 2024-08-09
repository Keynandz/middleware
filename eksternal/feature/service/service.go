package service

import (
	"github.com/labstack/echo/v4"
)

type Service interface {
	Get(c echo.Context) (string, error)
}
