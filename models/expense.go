package models

import "time"

// Expense define a estrutura de uma despesa
type Expense struct {
	ID          int
	Description string
	Amount      float64
	Date        time.Time
}

// NewExpense cria uma nova despesa
func NewExpense(id int, description string, amount float64) Expense {
	return Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}
}
