package routes

import (
	"easyfinance/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/expenses", controllers.AddExpense)
	r.GET("/expenses", controllers.GetExpenses)
	r.GET("/expenses/:id", controllers.GetExpenseByID)
	r.PUT("/expenses/:id", controllers.UpdateExpense)
	r.DELETE("/expenses/:id", controllers.DeleteExpense)

	return r
}
