package routes

import (
	"CodePonder/utils"
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

var SkipperRoutes = [4]string{"/login", "/signup", "/forgot-password", "/change-password"}

func IsAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if contains(SkipperRoutes, c.Path()) {
			return next(c)
		}

		userId, err := utils.IsLoggedIn(c)
		if err != nil {
			return c.JSON(err.Code, err.Message)
		}

		ctx := context.WithValue(c.Request().Context(), ContextUserKey, userId)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func CORSConfig() echo.MiddlewareFunc {
	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("CORS_ORIGIN")},
		AllowCredentials: true,
	})

	return cors
}

func contains(arr [4]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
