package main

import (
	"net/http"
	"os"

	"go-base-structure/pkg/db"
	"go-base-structure/pkg/minio"
	"go-base-structure/pkg/migration"
	"go-base-structure/pkg/util"
	"go-base-structure/pkg/util/env"
	"go-base-structure/routes"
    "go-base-structure/cron"

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

	// Minio
	minio.Init()

	// Migration
	migration.Init()

    // Cron
    cron.Init()

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

	if os.Getenv("SSL_CERT_PATH") != "" && os.Getenv("SSL_KEY_PATH") != "" {
		e.Logger.Fatal(e.StartTLS(":"+env.GetString("PORT"), os.Getenv("SSL_CERT_PATH"), os.Getenv("SSL_KEY_PATH")))
	} else {
		e.Logger.Fatal(e.Start(":" + env.GetString("PORT")))
	}
}
