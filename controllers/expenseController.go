package controllers

import (
	"database/sql"
	"easyfinance/database"
	"easyfinance/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	query := "INSERT INTO expenses (description, amount) VALUES ($1, $2) RETURNING id, date"
	err := database.DB.QueryRow(query, expense.Description, expense.Amount).Scan(&expense.ID, &expense.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar despesa"})
		return
	}

	c.JSON(http.StatusCreated, expense)
}

func GetExpenses(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, description, amount, date FROM expenses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar despesas"})
		return
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var exp models.Expense
		err := rows.Scan(&exp.ID, &exp.Description, &exp.Amount, &exp.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar dados"})
			return
		}
		expenses = append(expenses, exp)
	}

	c.JSON(http.StatusOK, expenses)
}

func GetExpenseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var expense models.Expense
	query := "SELECT id, description, amount, date FROM expenses WHERE id = $1"
	err = database.DB.QueryRow(query, id).Scan(&expense.ID, &expense.Description, &expense.Amount, &expense.Date)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Despesa não encontrada"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar despesa"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	query := "DELETE FROM expenses WHERE id = $1"
	res, err := database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir despesa"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Despesa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Despesa removida com sucesso"})
}

func UpdateExpense(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	query := "UPDATE expenses SET description = $1, amount = $2 WHERE id = $3"
	res, err := database.DB.Exec(query, expense.Description, expense.Amount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar despesa"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Despesa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Despesa atualizada com sucesso"})
}
