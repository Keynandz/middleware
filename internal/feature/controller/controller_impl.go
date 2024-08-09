package controller

import (
	"go-base-structure/eksternal/feature/service"
	res "go-base-structure/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl) Get(c echo.Context) error {
	result, err := co.Service.Get(c)
	if err != nil {
		return res.ErrorResponse(err)
	}

	return res.SuccessResponse(result).Send(c)
}
