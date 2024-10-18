package repository

import (
	"go-authorization/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	Find(c echo.Context, DB *gorm.DB, ID int) (entity.MasterRoleModel, error)
}
