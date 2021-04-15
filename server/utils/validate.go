package utils

import (
	"CodePonder/models"
	"strings"

	"github.com/labstack/echo/v4"
)

func ValidateRegister(user models.User) *echo.HTTPError {
	if len(user.Username) < 3 {
		return InvalidInput("username", "username must be at least 3 characters")
	}

	if strings.Contains(user.Username, "@") {
		return InvalidInput("username", "username connot contain @")
	}

	if !strings.Contains(user.Email, "@") || !strings.Contains(user.Email, ".") {
		return InvalidInput("email", "invalid email")
	}

	if len(user.Password) < 4 {
		return InvalidInput("password", "password must be at least 4 characts")
	}

	return nil
}

func InvalidInput(field, message string) *echo.HTTPError {
	return echo.NewHTTPError(400, models.ErrorMessage{
		Field:   field,
		Message: message,
	})
}
