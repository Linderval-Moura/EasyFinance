package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao verificar conexão com o banco:", err)
	}

	fmt.Println("Banco de dados conectado com sucesso!")
}
