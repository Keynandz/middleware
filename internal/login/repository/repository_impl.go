package repository

import (
	"go-authorization/model/dto"
	"go-authorization/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) Find(c echo.Context, DB *gorm.DB, payload dto.LoginRequest) (entity.MasterUserModel, error) {
	var result entity.MasterUserModel
	err := DB.Where("email = ?", payload.Email).First(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (r *repositoryImpl) FindListUser(c echo.Context, DB *gorm.DB) ([]entity.MasterUserModel, error) {
	var result []entity.MasterUserModel
	query := DB
	switch c.Get("role").(string) {
	case "Admin", "Super Admin", "Admin OPS", "GM":

	case "PIC", "Approver", "KAM", "BOD", "SO", "SPV Site", "GA", "SL", "SPV Area", "OH", "PM", "FSO":
		query = query.Where("isarea = ? AND isactive = 1", c.Get("isarea").(int))
	case "Employee":
		query = query.Where("id = ?", c.Get("id").(int))
	default:
		query = query.Where("working_area_id = ? AND isarea = ? AND isactive = 1", c.Get("working_area_id").(int), c.Get("isarea").(int))
	}
	err := query.Order("id ASC").Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, err
}
