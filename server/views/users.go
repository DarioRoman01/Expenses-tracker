package views

import (
	"CodePonder/cache"
	"CodePonder/controllers"
	"CodePonder/models"
	"CodePonder/utils"
	"context"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersHandler struct {
	DB *gorm.DB
}

var usersTable *controllers.UsersTable

func (u *UsersHandler) SignupView(c echo.Context) error {
	var userInput models.UserRegisterInput

	if err := c.Bind(&userInput); err != nil {
		return echo.NewHTTPError(423, "unable to parse request body")
	}

	isInvalid := utils.ValidateRegister(userInput)

	if isInvalid != nil {
		return c.JSON(isInvalid.Code, isInvalid.Message)
	}

	usersCreated, err := usersTable.CreateUser(&models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: userInput.Password,
	}, u.DB)

	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	session := cache.Default(c)
	session.Set("userId", usersCreated.ID)
	session.Save()
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

		session := cache.Default(c)
		session.Set("userId", user.ID)
		session.Save()

		return c.JSON(200, user)
	}
}

func (u *UsersHandler) ForgotPasswordView(c echo.Context) error {
	email := c.FormValue("email")
	user := usersTable.GetUserByEmail(email, u.DB)
	if user == nil {
		return utils.InvalidInput("email", "email does not exist")
	}

	ok := utils.SendEmail(user)

	if !ok {
		return echo.NewHTTPError(500, "unable to send email")
	}

	return c.JSON(200, "we send you and email please check your email")
}

func (u *UsersHandler) ChangePasswordView(c echo.Context) error {
	newPassword := c.FormValue("newPassword")
	if len(newPassword) < 4 {
		return utils.InvalidInput("newPassword", "password must be at least 4 characters")
	}

	ctx := context.Background()
	token := c.FormValue("token")
	redis := cache.Client()
	userId := redis.Get(ctx, fmt.Sprintf("forgot-password:%s", token))

	if userId.Val() == "" {
		return utils.InvalidInput("token", "invalid token")
	}

	user, err := usersTable.ChangePassword(userId.Val(), newPassword, u.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	session := cache.Default(c)
	session.Set("userId", user.ID)
	session.Save()

	return c.JSON(200, "password changed succesfully")
}
