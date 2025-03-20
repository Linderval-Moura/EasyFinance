package main

import (
	"fmt"
	"os"
	"strconv"

	"easyfinance/controllers"
	"easyfinance/services"
	"easyfinance/views"
)

func main() {
	for {
		fmt.Println("\nGerenciador de Despesas - EasyFinance")
		fmt.Println("1. Adicionar Despesa")
		fmt.Println("2. Atualizar Despesa")
		fmt.Println("3. Excluir Despesa")
		fmt.Println("4. Listar Despesas")
		fmt.Println("5. Resumo Geral")
		fmt.Println("6. Resumo Mensal")
		fmt.Println("7. Sair")
		fmt.Print("Escolha uma opção: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			var description string
			var amount float64
			fmt.Print("Descrição: ")
			fmt.Scanln(&description)
			fmt.Print("Valor: ")
			fmt.Scanln(&amount)

			if !services.ValidateExpense(amount) {
				views.ShowMessage("Erro: o valor da despesa deve ser positivo!")
				continue
			}
			expense := controllers.AddExpense(description, amount)
			views.ShowMessage(fmt.Sprintf("Despesa adicionada: %s - R$ %.2f", expense.Description, expense.Amount))

		case 2:
			var id int
			var description string
			var amount float64
			fmt.Print("ID da Despesa: ")
			fmt.Scanln(&id)
			fmt.Print("Nova Descrição: ")
			fmt.Scanln(&description)
			fmt.Print("Novo Valor: ")
			fmt.Scanln(&amount)

			if !services.ValidateExpense(amount) {
				views.ShowMessage("Erro: o valor da despesa deve ser positivo!")
				continue
			}

			err := controllers.UpdateExpense(id, description, amount)
			if err != nil {
				views.ShowMessage("Erro: " + err.Error())
			} else {
				views.ShowMessage("Despesa atualizada com sucesso!")
			}

		case 3:
			var id int
			fmt.Print("ID da Despesa: ")
			fmt.Scanln(&id)
			err := controllers.DeleteExpense(id)
			if err != nil {
				views.ShowMessage("Erro: " + err.Error())
			} else {
				views.ShowMessage("Despesa removida com sucesso!")
			}

		case 4:
			expenses := controllers.GetExpenses()
			views.ShowExpenses(expenses)

		case 5:
			total := controllers.GetTotalSummary()
			views.ShowSummary(total)

		case 6:
			var month int
			fmt.Print("Digite o número do mês (1-12): ")
			fmt.Scanln(&month)
			if month < 1 || month > 12 {
				views.ShowMessage("Mês inválido! Digite um número entre 1 e 12.")
			} else {
				total := controllers.GetMonthlySummary(month)
				views.ShowMonthlySummary(month, total)
			}

		case 7:
			views.ShowMessage("Saindo do programa...")
			os.Exit(0)

		default:
			views.ShowMessage("Opção inválida! Tente novamente.")
		}
	}
}
