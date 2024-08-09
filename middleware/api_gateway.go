package middleware

import (
	"net/http"
	"strings"
	res "go-base-structure/pkg/util/response"

	"github.com/labstack/echo/v4"
)

func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			err := "Token is missing"
			return res.CustomErrorBuilder(http.StatusUnauthorized, "something wrong", err).Send(c)
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		newToken, err := VerifyToken(tokenString)
		if err != nil {
			err := "Invalid token"
			return res.CustomErrorBuilder(http.StatusUnauthorized, "something wrong", err).Send(c)
		}

		if newToken != tokenString {
			c.Set("newToken", newToken)
		}

		return next(c)
	}
}