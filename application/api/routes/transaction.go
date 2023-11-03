package routes

import (
	"github.com/jvfrodrigues/transaction-product-wex/application/api/controllers"
	"github.com/jvfrodrigues/transaction-product-wex/application/treasury"
	"github.com/jvfrodrigues/transaction-product-wex/infra/repository"
)

func (r *Routes) setupTransactionRoutes() {
	exchangeService := treasury.NewTreasuryExchange()
	transactionRepository := repository.NewTrasactionRepositoryDb(r.Database)
	transactionController := controllers.NewTransactionController(transactionRepository, exchangeService)
	transactionGroup := r.Router.Group("/transaction")
	{
		transactionGroup.POST("/", transactionController.RegisterTransaction)
		transactionGroup.GET("/exchange/:id", transactionController.FindTransactionAndExchangeCurrency)
	}
}
