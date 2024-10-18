package repository

import (
	"go-authorization/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) Find(c echo.Context, DB *gorm.DB, ID int) (entity.MasterRoleModel, error) {
	var result entity.MasterRoleModel
	err := DB.First(&result, ID).Error
	if err != nil {
		return result, err
	}
	return result, err
}
