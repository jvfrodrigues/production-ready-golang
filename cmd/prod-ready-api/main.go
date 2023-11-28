package main

import (
	"os"

	"github.com/jvfrodrigues/production-ready-golang/internal/application/api"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/db"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/logger/zap"
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
		port = "8888"
	}
	serverPort := ":" + port
	return serverPort
}
