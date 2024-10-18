package main

import (
	"net/http"
	"os"

	"go-authorization/pkg/db"
	"go-authorization/pkg/migration"
	"go-authorization/pkg/util"
	"go-authorization/pkg/util/env"
	"go-authorization/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()

	// Env
	env := env.NewEnv()
	env.Load()

	// Database
	db.Init()

	// Migration
	migration.Init()

	// Validator
	e.Validator = &util.CustomValidation{Validator: validator.New()}
	logrus.SetOutput(os.Stdout)

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Auth-Token"},
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	// Route
	routes.Init(e.Group("/api/v1"))
	routes.Interface(e.Group(""), e)

	if os.Getenv("SSL_CERT_PATH") != "" && os.Getenv("SSL_KEY_PATH") != "" {
		e.Logger.Fatal(e.StartTLS(":"+env.GetString("PORT"), os.Getenv("SSL_CERT_PATH"), os.Getenv("SSL_KEY_PATH")))
	} else {
		e.Logger.Fatal(e.Start(":" + env.GetString("PORT")))
	}
}
