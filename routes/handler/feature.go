package handler

import (
	"go-base-structure/internal/feature/controller"
	"go-base-structure/internal/feature/repository"
	"go-base-structure/internal/feature/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handlerFeature struct {
	Controller controller.Controller
}

func NewFeatureHandler(db *gorm.DB) *handlerFeature {
	sr := repository.NewRepository()
	ss := service.NewService(db, sr)

	return &handlerFeature{
		Controller: controller.NewController(ss),
	}
}

func (h *handlerFeature) Route(g *echo.Group) {
	g.GET("", h.Controller.Get)
}