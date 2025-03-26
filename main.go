package main

import (
	//	"easyfinance/database"
	"easyfinance/routes"
	"fmt"
)

func main() {
	//	database.Connect()
	r := routes.SetupRouter()
	fmt.Println("Servidor rodando na porta 8080...")
	r.Run(":8080")
}
