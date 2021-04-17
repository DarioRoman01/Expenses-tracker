package views

import (
	"CodePonder/controllers"
	"CodePonder/models"
	"CodePonder/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ExpensesViews struct {
	DB *gorm.DB
}

var expensesTable *controllers.ExpensesTable

func (e *ExpensesViews) AddExpenseView(c echo.Context) error {
	userId, err := utils.IsLoggedIn(c)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	var expense models.Expenses
	if err := c.Bind(&expense); err != nil {
		return echo.NewHTTPError(423, "cannot parse request body")
	}

	expense.UserID = userId
	expenseCreated, err := expensesTable.CreateExpense(&expense, e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(201, expenseCreated)
}

func (e *ExpensesViews) ExpensesView(c echo.Context) error {
	userId, httpErr := utils.IsLoggedIn(c)
	if httpErr != nil {
		return c.JSON(httpErr.Code, httpErr.Message)
	}

	limit := c.QueryParam("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(400, "invalid limit")
	}

	cursor := c.QueryParam("cursor")
	expenses, httpErr, hasMore := expensesTable.GetUserExpenses(userId, intLimit, &cursor, e.DB)
	if httpErr != nil {
		return c.JSON(httpErr.Code, httpErr.Message)
	}

	return c.JSON(200, &models.PaginatedExpenses{
		Expenses: expenses,
		HasMore:  hasMore,
	})
}

func (e *ExpensesViews) ExpensesByCategoryView(c echo.Context) error {
	userId, err := utils.IsLoggedIn(c)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	category := c.QueryParam("category")
	expenses, err := expensesTable.GetExpenseByCategory(userId, category, e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(200, expenses)
}
