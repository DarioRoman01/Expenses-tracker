package controllers

import (
	"CodePonder/models"
	"CodePonder/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var passwordCfg *utils.PasswordConfig

type UsersTable struct{}

func init() {
	passwordCfg = &utils.PasswordConfig{
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}
}

func (u *UsersTable) CreateUser(user *models.User, db *gorm.DB) (*models.User, *echo.HTTPError) {
	usernameTaken := u.getUserByUsername(user.Username, db)
	if usernameTaken != nil {
		return nil, utils.InvalidInput("username", "username already in use")
	}

	emailTaken := u.getUserByEmail(user.Email, db)
	if emailTaken != nil {
		return nil, utils.InvalidInput("email", "email already in use")
	}

	hashPassword, err := utils.GeneratePassword(passwordCfg, user.Password)
	if err != nil {
		return nil, echo.NewHTTPError(500, "unable to hash password")
	}

	user.Password = hashPassword
	if err := db.Create(&user).Error; err != nil {
		return nil, echo.NewHTTPError(500, "unable to create user")
	}

	return user, nil
}

func (u *UsersTable) LoginByEmail(input *models.UserInput, db *gorm.DB) (*models.User, *echo.HTTPError) {
	user := u.getUserByEmail(input.UsernameOrEmail, db)
	if user == nil {
		return nil, utils.InvalidInput("email", "email does not exist")
	}

	ok, _ := utils.ComparePasswords(input.Password, user.Password)
	if !ok {
		return nil, utils.InvalidInput("password", "invalid credentials")
	}

	return user, nil
}

func (u *UsersTable) LoginByUsername(input *models.UserInput, db *gorm.DB) (*models.User, *echo.HTTPError) {
	user := u.getUserByUsername(input.UsernameOrEmail, db)
	if user == nil {
		return nil, utils.InvalidInput("username", "username does not exist")
	}

	ok, _ := utils.ComparePasswords(input.Password, user.Password)

	if !ok {
		return nil, utils.InvalidInput("password", "invalid credentials")
	}

	return user, nil
}

func (u *UsersTable) getUserByUsername(username string, db *gorm.DB) *models.User {
	var user models.User

	db.Find(&user).Where("username = ?", username)

	if user.ID != 0 {
		return nil
	}

	return &user
}

func (u *UsersTable) getUserByEmail(email string, db *gorm.DB) *models.User {
	var user models.User

	db.Find(&user).Where("email = ?", email)

	if user.ID != 0 {
		return nil
	}

	return &user
}
