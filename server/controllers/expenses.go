package controllers

import (
	"CodePonder/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ExpensesTable struct{}

func (e *ExpensesTable) CreateExpense(expense *models.Expenses, db *gorm.DB) (*models.Expenses, *echo.HTTPError) {
	if err := db.Create(&expense).Error; err != nil {
		return nil, echo.NewHTTPError(500, "unable to create expense")
	}

	return expense, nil
}

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

func (e *ExpensesTable) GetUserExpenses(userId int, limit int, cursor *string, db *gorm.DB) ([]models.Expenses, *echo.HTTPError) {
	var expenses []models.Expenses
	if limit > 50 {
		limit = 50
	}

	if cursor != nil {
		db.Raw(`
			SELECT * FROM "expenses"
			WHERE user_id = ?
			WHERE created_at < ?
			ORBER BY created_at DESC
			LIMIT ?
		`, userId, cursor, limit).Find(&expenses)
	} else {
		db.Raw(`
			SELECT * FROM "expenses"
			WHERE user_id = ?
			ORBER BY created_at DESC
			LIMIT ?
		`, userId, limit).Find(&expenses)
	}

	return expenses, nil
}

func (e *ExpensesTable) GetExpenseByCategory(category string, db *gorm.DB) ([]models.Expenses, *echo.HTTPError) {
	var expenses []models.Expenses

	db.Model(&models.Expenses{}).Where("category = ?", category).Find(&expenses)

	if len(expenses) == 0 {
		return nil, echo.NewHTTPError(404, "invalid category")
	}

	return expenses, nil
}

func (e *ExpensesTable) CreateCategory(category *models.Category, db *gorm.DB) (*models.Category, *echo.HTTPError) {
	if err := db.Create(&category).Error; err != nil {
		return nil, echo.NewHTTPError(500, "unable to create expense")
	}

	return category, nil
}
