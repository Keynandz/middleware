package service

import (
	"go-authorization/model/dto"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Get(c echo.Context, payload dto.LoginRequest) (dto.UserResponse, error)
	GetUser(c echo.Context) ([]dto.ListUserResponse, error)
}
