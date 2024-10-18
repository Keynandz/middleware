package handler

import (
	roleRepository "go-authorization/internal/role/repository"
	token "go-authorization/middleware"
	"go-authorization/internal/login/controller"
	"go-authorization/internal/login/repository"
	"go-authorization/internal/login/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handlerFeature struct {
	Controller controller.Controller
}

func NewFeatureHandler(db *gorm.DB) *handlerFeature {
	sr := repository.NewRepository()
	rr := roleRepository.NewRepository()
	ss := service.NewService(db, sr, rr)

	return &handlerFeature{
		Controller: controller.NewController(ss),
	}
}

func (h *handlerFeature) Route(g *echo.Group) {
	authGroup := g.Group("")
	authGroup.Use(token.JWTAuthMiddleware)

	g.POST("/login", h.Controller.Login)
	authGroup.GET("/user", h.Controller.Get)
}