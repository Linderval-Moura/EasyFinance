package controllers

import (
	"easyfinance/database"
	"easyfinance/models"
)

func AddExpense(description string, amount float64) models.Expense {
	return database.AddExpense(description, amount)
}

func GetExpenses() []models.Expense {
	return database.GetExpenses()
}

func DeleteExpense(id int) error {
	return database.DeleteExpense(id)
}

func UpdateExpense(id int, description string, amount float64) error {
	return database.UpdateExpense(id, description, amount)
}

func GetMonthlySummary(month int) float64 {
	return database.GetMonthlySummary(month)
}

func GetTotalSummary() float64 {
	return database.GetTotalSummary()
}
