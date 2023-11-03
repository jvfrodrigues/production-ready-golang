package main

import (
	"os"

	"github.com/jvfrodrigues/transaction-product-wex/application/api"
	"github.com/jvfrodrigues/transaction-product-wex/infra/db"
	"github.com/jvfrodrigues/transaction-product-wex/infra/logger/zap"
)

func main() {
	logger := zap.NewLogger()
	database := db.ConnectDB(os.Getenv("env"))
	server := api.NewServer("release")
	port := preparePort()
	logger.Info("Server running on " + port)
	server.StartServer(database, port)
}

func preparePort() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	serverPort := ":" + port
	return serverPort
}
