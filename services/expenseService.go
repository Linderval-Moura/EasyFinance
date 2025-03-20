package services

import "easyfinance/models"

// Validar se a despesa tem valor válido
func ValidateExpense(amount float64) bool {
	return amount > 0
}
