package views

import (
	"CodePonder/controllers"
	"CodePonder/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ExpensesViews struct {
	DB *gorm.DB
}

var expensesTable *controllers.ExpensesTable

func (e *ExpensesViews) AddExpenseView(c echo.Context) error {
	userId := c.Request().Context().Value("user").(int)

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
	userId := c.Request().Context().Value("user").(int)
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
	userId := c.Request().Context().Value("user").(int)

	category := c.QueryParam("category")
	expenses, err := expensesTable.GetExpenseByCategory(userId, category, e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(200, expenses)
}

func (e *ExpensesViews) DeleteExpenseView(c echo.Context) error {
	userId := c.Request().Context().Value("user").(int)
	err := expensesTable.DeleteExpense(c.Param("id"), userId, e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(200, "succesfully deleted")
}

func (e *ExpensesViews) GetCategorysView(c echo.Context) error {
	categorys, err := expensesTable.GetCategorys(e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(200, categorys)
}

func (e *ExpensesViews) AvarageAmountView(c echo.Context) error {
	userId := c.Request().Context().Value("user").(int)

	avarage, err := expensesTable.GetAvarageAmount(userId, e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(200, avarage)
}

func (e *ExpensesViews) UpdateExpenseView(c echo.Context) error {
	userId := c.Request().Context().Value("user").(int)

	var data models.Expenses
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(423, "unable to parse request body")
	}

	expense, err := expensesTable.ModifyExpense(userId, c.Param("id"), data, e.DB)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}

	return c.JSON(200, expense)
}
