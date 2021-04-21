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

	psql.AutoMigrate(&models.User{}, &models.Category{}, &models.Expenses{})
	usersViews := views.UsersViews{DB: psql}
	expensesviews := views.ExpensesViews{DB: psql}

	// users views
	e.POST("/signup", usersViews.SignupView)
	e.POST("/login", usersViews.LoginView)
	e.GET("/me", usersViews.MeView)
	e.POST("logout", usersViews.LogoutView)
	e.POST("/forgot-password", usersViews.ForgotPasswordView)
	e.POST("/change-password", usersViews.ChangePasswordView)

	// Expenses views
	e.GET("/expenses", expensesviews.ExpensesView)
	e.GET("/expenses/:category", expensesviews.ExpensesByCategoryView)
	e.POST("/expenses", expensesviews.AddExpenseView)
	e.GET("/categorys", expensesviews.GetCategorysView)
	e.GET("/categorys/avarage", expensesviews.AvarageAmountView)
	e.DELETE("/expenses/:id", expensesviews.DeleteExpenseView)
	e.PATCH("expenses/:id", expensesviews.UpdateExpenseView)
}
