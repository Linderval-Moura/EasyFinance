# EasyFinance
Rastreador de despesas para gerenciar finanças. Permite que os usuários adicionem, excluam e visualizem suas despesas. Também fornecerá um resumo das despesas.

## Estrutura Inicial do Projeto

### Backend

```
easyfinance/
│── main.go
│── models/
│   ├── expense.go
│── controllers/
│   ├── expenseController.go
│── services/
│   ├── expenseService.go
│── views/
│   ├── cliView.go
│── database/
│   ├── db.go
```

## Como Executar
1. Baixe o código e crie a estrutura de diretórios conforme o modelo.
2. Vá até a pasta do projeto e execute:
```
go run main.go
```
Siga as opções no menu para adicionar, atualizar e visualizar despesas.