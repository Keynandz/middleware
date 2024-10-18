package repository

import (
	"go-authorization/model/dto"
	"go-authorization/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	Find(c echo.Context, DB *gorm.DB, payload dto.LoginRequest) (entity.MasterUserModel, error)
	FindListUser(c echo.Context, DB *gorm.DB) ([]entity.MasterUserModel, error)
}