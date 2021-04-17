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

// Create user validates that the username  or email is already in use and hash users password
func (u *UsersTable) CreateUser(user *models.User, db *gorm.DB) (*models.User, *echo.HTTPError) {
	usernameTaken := u.getUserByUsername(user.Username, db)
	if usernameTaken != nil {
		return nil, utils.InvalidInput("username", "username already in use")
	}

	emailTaken := u.GetUserByEmail(user.Email, db)
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

// Login by email handles users login with email and validate that given email exist
func (u *UsersTable) LoginByEmail(input *models.UserInput, db *gorm.DB) (*models.User, *echo.HTTPError) {
	user := u.GetUserByEmail(input.UsernameOrEmail, db)
	if user == nil {
		return nil, utils.InvalidInput("email", "email does not exist")
	}

	ok, _ := utils.ComparePasswords(input.Password, user.Password)
	if !ok {
		return nil, utils.InvalidInput("password", "invalid credentials")
	}

	return user, nil
}

// Login by username handles users login with username and validate that given username exist
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

// Change Password change user password
func (u *UsersTable) ChangePassword(id, newPassword string, db *gorm.DB) (*models.User, *echo.HTTPError) {
	user := u.GetUserByid(id, db)
	if user == nil {
		return nil, utils.InvalidInput("token", "invalid token")
	}

	hashPwd, _ := utils.GeneratePassword(passwordCfg, newPassword)
	db.Model(&user).Update("password", hashPwd)

	return user, nil
}

// retrieve user by username
func (u *UsersTable) getUserByUsername(username string, db *gorm.DB) *models.User {
	var user models.User
	db.Model(&models.User{}).Where("username = ?", username).Find(&user)
	if user.ID == 0 {
		return nil
	}

	return &user
}

// return user by email
func (u *UsersTable) GetUserByEmail(email string, db *gorm.DB) *models.User {
	var user models.User
	db.Model(&models.User{}).Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return nil
	}

	return &user
}

// return user by id
func (u *UsersTable) GetUserByid(id interface{}, db *gorm.DB) *models.User {
	var user models.User
	db.First(&user, id)
	if user.ID == 0 {
		return nil
	}

	return &user
}
