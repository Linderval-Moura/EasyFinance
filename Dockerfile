FROM golang:1.20-alpine

WORKDIR /app

# Copiando o go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Dependências
RUN go mod tidy

COPY . .

# Pacote para carregar arquivo .env
RUN go get github.com/joho/godotenv

RUN go build -o easyfinance .

EXPOSE 8080

# Comando para rodar o aplicativo
CMD ["./easyfinance"]