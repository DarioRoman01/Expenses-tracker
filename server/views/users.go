package views

import (
	"CodePonder/controllers"
	"CodePonder/models"
	"CodePonder/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersHandler struct {
	DB *gorm.DB
}

var usersTable *controllers.UsersTable

func (u *UsersHandler) SignupView(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(423, "unable to parse request body")
	}

	isInvalid := utils.ValidateRegister(user)

	if isInvalid != nil {
		return c.JSON(isInvalid.Code, isInvalid.Message)
	}

	usersCreated, err := usersTable.CreateUser(&user, u.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(201, usersCreated)
}

func (u *UsersHandler) LoginView(c echo.Context) error {
	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		return echo.NewHTTPError(423, "unable to parse request body")
	}

	if strings.Contains(userInput.UsernameOrEmail, "@") {
		user, err := usersTable.LoginByEmail(&userInput, u.DB)
		if err != nil {
			return c.JSON(err.Code, err.Message)
		}

		return c.JSON(200, user)

	} else {
		user, err := usersTable.LoginByUsername(&userInput, u.DB)
		if err != nil {
			return c.JSON(err.Code, err.Message)
		}

		return c.JSON(200, user)
	}
}
