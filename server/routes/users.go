package routes

import (
	"CodePonder/db"
	"CodePonder/models"
	"CodePonder/views"
	"log"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	psql, err := db.Connect()
	if err != nil {
		log.Fatal("unable to connect to postgres: ", err)
	}

	psql.AutoMigrate(&models.User{})
	usersViews := views.UsersViews{DB: psql}

	// users views
	e.POST("/login", usersViews.LoginView)
	e.POST("/signup", usersViews.SignupView)
	e.POST("/forgot-password", usersViews.ForgotPasswordView)
	e.POST("/change-password", usersViews.ChangePasswordView)
	e.GET("/me", usersViews.MeView)
}
