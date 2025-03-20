package views

import (
	"fmt"
	"easyfinance/models"
)

// ShowExpenses exibe todas as despesas de maneira organizada
func ShowExpenses(expenses []models.Expense) {
	fmt.Println("\nLista de Despesas:")
	for _, expense := range expenses {
		fmt.Printf("ID: %d | Descrição: %s | Valor: R$ %.2f | Data: %s\n",
			expense.ID, expense.Description, expense.Amount, expense.Date.Format("02/01/2006"))
	}
}

// ShowSummary exibe o resumo geral das despesas
func ShowSummary(total float64) {
	fmt.Printf("\nResumo das despesas: R$ %.2f\n", total)
}

// ShowMonthlySummary exibe o resumo das despesas de um mês específico
func ShowMonthlySummary(month int, total float64) {
	fmt.Printf("\nResumo do mês %d: R$ %.2f\n", month, total)
}

// ShowMessage exibe uma mensagem ao usuário
func ShowMessage(message string) {
	fmt.Println(message)
}
