# EasyFinance
API para gerenciar finanças. Permite que os usuários adicionem, excluam e visualizem suas despesas. Também fornecerá um resumo das despesas.

## Estrutura Inicial do Projeto

### Backend

```
easyfinance/
│── main.go
│── models/
│   ├── expense.go
│── database/
│   ├── db.go
│── controllers/
│   ├── expenseController.go
│── routes/
│   ├── expenseRoutes.go
│── go.mod
│── go.sum

```

## Como Executar
1. Baixe o código e crie a estrutura de diretórios conforme o modelo.
2. Vá até a pasta do projeto e execute:
```
go run main.go
```
Siga as opções no menu para adicionar, atualizar e visualizar despesas.