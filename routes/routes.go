package routes

import (
	"go-base-structure/pkg/constant"
	"go-base-structure/pkg/db"
	"go-base-structure/pkg/util/env"
	"go-base-structure/routes/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	db, err := db.GetConnection(constant.DB_BASE_STRUCTURE)
	if err != nil {
		panic("Failed init db, connection is undefined")
	}

	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+env.NewEnv().GetString("APP")+"! version "+env.NewEnv().GetString("VERSION")+" in mode "+env.NewEnv().GetString("ENV"))
	})

	// Routes
	handler.NewFeatureHandler(db).Route(g.Group("/feature"))
}
