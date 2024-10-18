package routes

import (
	"go-authorization/pkg/constant"
	"go-authorization/pkg/db"
	"go-authorization/pkg/util/env"
	"go-authorization/routes/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	db, err := db.GetConnection(constant.DB_BASE_STRUCTURE)
	if err != nil {
		panic("Failed init db, connection is undefined")
	}

	// Route welcome message
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+env.NewEnv().GetString("APP")+"! version "+env.NewEnv().GetString("VERSION")+" in mode "+env.NewEnv().GetString("ENV"))
	})

	// Route for /feature
	handler.NewFeatureHandler(db).Route(g.Group("/feature"))
}

func Interface(g *echo.Group, e *echo.Echo) {
	e.Static("/static", "interface")

	g.GET("/login", func(c echo.Context) error {
		return c.File("interface/login.html")
	})

	g.GET("/index", func(c echo.Context) error {
		return c.File("interface/index.html")
	})
}
