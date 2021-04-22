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

type UsersViews struct {
	DB *gorm.DB
}

var usersTable *controllers.UsersTable

//Users singup view. Validates users input
//and create a session when the account is created
func (u *UsersViews) SignupView(c echo.Context) error {
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

//Users Login view. check is user in is loggin with
//email or username and generate cookie session
func (u *UsersViews) LoginView(c echo.Context) error {
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

// Forgot Password View check if the email is in use and
// send an email to users email
func (u *UsersViews) ForgotPasswordView(c echo.Context) error {
	var email map[string]string
	err := (&echo.DefaultBinder{}).BindBody(c, &email)
	if err != nil {
		return c.JSON(423, "unable to parse request body")
	}
	user := usersTable.GetUserByEmail(email["email"], u.DB)
	if user == nil {
		return utils.InvalidInput("email", "email does not exist")
	}

	ok := utils.SendEmail(user)

	if !ok {
		return echo.NewHTTPError(500, "unable to send email")
	}

	return c.JSON(200, "we send you and email please check your email")
}

// Change password view validate if given token is store in redis
// and validate that the token is valid and the new password
func (u *UsersViews) ChangePasswordView(c echo.Context) error {
	var body map[string]string
	err := (&echo.DefaultBinder{}).BindBody(c, &body)
	if err != nil {
		return c.JSON(400, "unable to parse request body")
	}

	if len(body["newPassword"]) < 4 {
		return utils.InvalidInput("newPassword", "password must be at least 4 characters")
	}

	ctx := context.Background()
	redis := cache.Client()
	key := fmt.Sprintf("forgot-password:%s", body["token"])
	userId := redis.Get(ctx, key)

	if userId.Val() == "" {
		return utils.InvalidInput("token", "invalid token")
	}

	user, httpErr := usersTable.ChangePassword(userId.Val(), body["newPassword"], u.DB)
	if err != nil {
		return c.JSON(httpErr.Code, httpErr.Message)
	}

	session := cache.Default(c)
	session.Set("userId", user.ID)
	session.Save()
	redis.Del(ctx, key)
	return c.JSON(200, "password changed succesfully")
}

// me view return users info when is logged in
func (u *UsersViews) MeView(c echo.Context) error {
	session := cache.Default(c)
	val := session.Get("userId")
	if val == nil {
		return echo.NewHTTPError(401, "not authenticated")
	}

	user := usersTable.GetUserByid(val, u.DB)
	return c.JSON(200, user)
}

// Logout view handles users logout destroing sessions
func (u *UsersViews) LogoutView(c echo.Context) error {
	session := cache.Default(c)
	session.Delete("userId")
	session.Clear()
	session.Save()
	return c.JSON(200, true)
}
