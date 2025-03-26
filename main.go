package main

import (
	//	"easyfinance/database"
	"easyfinance/routes"
	"fmt"
	"os"
)

func main() {
	//	database.Connect()
	port := os.Getenv("PORT") // Busca a porta das variáveis de ambiente
	if port == "" {
		fmt.Println("Environment PORT não declarada!") // Usa 8080 como padrão se nenhuma porta for especificada
	}

	route := routes.SetupRouter()
	fmt.Println("Servidor rodando...")
	route.Run(":" + port)
}
