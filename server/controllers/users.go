package controllers

import (
	"CodePonder/models"
	"CodePonder/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var passwordCfg *utils.PasswordConfig

type UsersHandler struct {
	DB *gorm.DB
}

func init() {
	passwordCfg = &utils.PasswordConfig{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
}

func (u *UsersHandler) CreateUser(user models.User) (*models.User, *echo.HTTPError) {
	usernameTaken := u.getUserByUsername(user.Username)
	if usernameTaken != nil {
		return nil, echo.NewHTTPError(400, models.ErrorMessage{
			Field:   "username",
			Message: "username already taken",
		})
	}

	emailTaken := u.getUserByEmail(user.Email)
	if emailTaken != nil {
		return nil, echo.NewHTTPError(400, models.ErrorMessage{
			Field:   "email",
			Message: "email already in use",
		})
	}

	hashPassword, err := utils.GeneratePassword(passwordCfg, user.Password)
	if err != nil {
		return nil, echo.NewHTTPError(500, "unable to hash password")
	}

	user.Password = hashPassword
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, echo.NewHTTPError(500, "unable to create user")
	}

	return &user, nil
}

func (u *UsersHandler) getUserByUsername(username string) *models.User {
	var user models.User

	u.DB.Find(&user).Where("username = ?", username)

	if user.ID != 0 {
		return nil
	}

	return &user
}

func (u *UsersHandler) getUserByEmail(email string) *models.User {
	var user models.User

	u.DB.Find(&user).Where("email = ?", email)

	if user.ID != 0 {
		return nil
	}

	return &user
}
