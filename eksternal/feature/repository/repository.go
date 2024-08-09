package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	Find(c echo.Context, DB *gorm.DB) (string, error)
}