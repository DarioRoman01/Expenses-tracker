package utils

import (
	"CodePonder/cache"
	"CodePonder/models"
	"strings"

	"github.com/labstack/echo/v4"
)

// Validate users input when they are register
func ValidateRegister(input models.UserRegisterInput) *echo.HTTPError {
	if len(input.Username) < 3 {
		return InvalidInput("username", "username must be at least 3 characters")
	}

	if strings.Contains(input.Username, "@") {
		return InvalidInput("username", "username connot contain @")
	}

	if !strings.Contains(input.Email, "@") || !strings.Contains(input.Email, ".") {
		return InvalidInput("email", "invalid email")
	}

	if len(input.Password) < 4 {
		return InvalidInput("password", "password must be at least 4 characts")
	}

	return nil
}

// return a bad request error with a field and
// message for the error.
func InvalidInput(field, message string) *echo.HTTPError {
	return echo.NewHTTPError(400, models.ErrorMessage{
		Field:   field,
		Message: message,
	})
}

// validate if user is logged in
func IsLoggedIn(c echo.Context) (int, *echo.HTTPError) {
	session := cache.Default(c)
	userId := session.Get("userId")
	if userId == "" || userId == nil {
		return 0, echo.NewHTTPError(401, "not authenticated")
	}

	return userId.(int), nil
}
