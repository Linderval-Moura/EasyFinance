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

## Como Executar 🚀
1. Instale as dependências:
```
go mod tidy
```
2. Execute a API:
```
go run main.go
```
3. Teste os endpoints via Postman ou cURL.

