package main

import (
	"os"

	"github.com/jvfrodrigues/transaction-product-wex/application/api"
	"github.com/jvfrodrigues/transaction-product-wex/infra/db"
)

func main() {
	database := db.ConnectDB(os.Getenv("env"))
	server := api.NewServer("release")
	server.StartServer(database, ":8080")
}
