package main

import (
	//	"easyfinance/database"
	"github.com/joho/godotenv"
	"easyfinance/routes"
	"fmt"
	"os"
)

func main() {
	//	database.Connect()
	port := os.Getenv("PORT") // Busca a porta das variáveis de ambiente
	if port == "" {
		fmt.Println("Environment PORT não declarada!")
		port = ("8080")
		fmt.Println("Servidor rodando na porta 8080...") // Usa 8080 como padrão se nenhuma porta for especificada 		
	}

	route := routes.SetupRouter()
	fmt.Println("Servidor rodando...")
	route.Run(":" + port)
}
