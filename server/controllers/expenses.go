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

func (e *ExpensesTable) GetExpense(id string, userId int, db *gorm.DB) (*models.Expenses, *echo.HTTPError) {
	var expense models.Expenses

	db.Model(&models.Expenses{}).Where("id = ?", id).Find(&expense)

	if expense.ID == 0 {
		return nil, echo.NewHTTPError(404, "expense not found")
	}

	if expense.UserID != userId {
		return nil, echo.NewHTTPError(404, "you do not have permissions to perfom this action")
	}

	return &expense, nil
}

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
			WHERE created_at < ?
			ORBER BY created_at DESC
			LIMIT ?
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
		return nil, echo.NewHTTPError(400, "invalid limit"), false
	}

	if len(expenses) == limit {
		return expenses[0 : limit-1], nil, true
	}

	return expenses[0 : len(expenses)-1], nil, false
}

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

func (e *ExpensesTable) GetCategorys(db *gorm.DB) ([]*models.Category, *echo.HTTPError) {
	var categorys []*models.Category

	db.Find(&categorys)

	if len(categorys) == 0 {
		return nil, echo.NewHTTPError(500, "somethin wrong happend :(")
	}

	return categorys, nil
}

func (e *ExpensesTable) GetAvarageAmount(userId int, db *gorm.DB) (float64, *echo.HTTPError) {
	var avarage float64

	db.Raw(`
	SELECT AVG(amount) FROM expenses
	WHERE user_id = ?
	`, userId).Find(&avarage)

	if avarage == 0 {
		return 0, echo.NewHTTPError(500, "something wrong happend")
	}

	return avarage, nil
}

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
