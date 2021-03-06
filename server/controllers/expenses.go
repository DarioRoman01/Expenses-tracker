package controllers

import (
	"CodePonder/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ExpensesTable struct{}

// CreateExpense handle expenses creation in the db
func (e *ExpensesTable) CreateExpense(expense *models.Expenses, db *gorm.DB) (*models.Expenses, *echo.HTTPError) {
	if err := db.Create(&expense).Error; err != nil {
		return nil, echo.NewHTTPError(500, "unable to create expense")
	}

	return expense, nil
}

// DeleteExpense delete expense from the db and validate
// that the requesting user is owner of the expense
func (e *ExpensesTable) DeleteExpense(id string, userId int, db *gorm.DB) *echo.HTTPError {
	var expense models.Expenses
	db.First(&expense, id)
	if expense.ID == 0 {
		return echo.NewHTTPError(404, "expense not found")
	}

	if expense.UserID != userId {
		return echo.NewHTTPError(403, "you do not have permission to perform this action")
	}

	db.Delete(&expense)

	return nil
}

// GetExpense retrieve expense by id and validate
// that the requesting user is owner of the expense
func (e *ExpensesTable) GetExpense(id string, userId int, db *gorm.DB) (*models.Expenses, *echo.HTTPError) {
	var expense models.Expenses

	db.Model(&models.Expenses{}).Where("id = ?", id).Find(&expense)

	if expense.ID == 0 {
		return nil, echo.NewHTTPError(404, "expense not found")
	}

	if expense.UserID != userId {
		return nil, echo.NewHTTPError(403, "you do not have permissions to perfom this action")
	}

	return &expense, nil
}

// retrieve users expenses using limit and cursor
func (e *ExpensesTable) GetUserExpenses(userId int, limit int, cursor *string, db *gorm.DB) ([]models.Expenses, *echo.HTTPError, bool) {
	var expenses []models.Expenses
	if limit > 50 {
		limit = 50
	}
	limit++

	if cursor != nil && *cursor != "" {
		db.Raw(`
			SELECT e.*
			FROM expenses e
			WHERE user_id = ?
			AND created_at < ?
			ORDER BY created_at DESC
			LIMIT ?;
		`, userId, cursor, limit).Find(&expenses)
	} else {
		db.Raw(`
			select e.*
			FROM expenses e
			WHERE user_id = ?
			ORDER BY created_at DESC
			LIMIT ?;
		`, userId, limit).Find(&expenses)
	}

	if len(expenses) == 0 {
		return expenses, nil, false
	}

	if len(expenses) == limit {
		return expenses[0 : limit-1], nil, true
	}

	return expenses, nil, false
}

// retrieve expenses by category
func (e *ExpensesTable) GetExpenseByCategory(userId int, category string, db *gorm.DB) ([]models.Expenses, *echo.HTTPError) {
	var expenses []models.Expenses

	db.Raw(`
		SELECT * FROM "expenses"
		WHERE user_id = ?
		WHERE category_name = ?
		ORDER BY created_at DESC
	`, userId, category).Find(&expenses)

	if len(expenses) == 0 {
		return nil, echo.NewHTTPError(404, "invalid category")
	}

	return expenses, nil
}

// list all categorys
func (e *ExpensesTable) GetCategorys(db *gorm.DB) ([]*models.Category, *echo.HTTPError) {
	var categorys []*models.Category

	db.Find(&categorys)

	if len(categorys) == 0 {
		return nil, echo.NewHTTPError(500, "somethin wrong happend :(")
	}

	return categorys, nil
}

// get avarage amount of expenses of the user
func (e *ExpensesTable) GetAvarageAmount(userId int, db *gorm.DB) (float64, *echo.HTTPError) {
	var avarage float64

	err := db.Raw(`
	SELECT AVG(amount) FROM expenses
	WHERE user_id = ?
	`, userId).Find(&avarage).Error

	if err != nil {
		return 0, nil
	}

	return avarage, nil
}

// handle expoenses updates and validate
// that the requesting user in owner of the expense
func (e *ExpensesTable) ModifyExpense(userId int, id string, data models.Expenses, db *gorm.DB) (*models.Expenses, *echo.HTTPError) {
	var storeExpense models.Expenses

	db.Model(&models.Expenses{}).Where("id = ?", id).Find(&storeExpense)
	if storeExpense.ID == 0 {
		return nil, echo.NewHTTPError(404, "post not found")
	}

	if storeExpense.UserID != userId {
		return nil, echo.NewHTTPError(403, "you do not have permissions to perform this action")
	}

	if err := db.Model(&storeExpense).Updates(data).Error; err != nil {
		return nil, echo.NewHTTPError(500, "unable to update expense")
	}

	return &storeExpense, nil
}
