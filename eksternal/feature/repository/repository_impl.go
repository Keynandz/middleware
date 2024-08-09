package repository

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) Find(c echo.Context, DB *gorm.DB) (string, error) {
	result := "Contoh Data Ges"
	return result, nil
}
