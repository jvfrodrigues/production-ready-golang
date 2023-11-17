package routes

import (
	"github.com/jvfrodrigues/production-ready-golang/internal/application/api/controllers"
	"github.com/jvfrodrigues/production-ready-golang/internal/application/treasury"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/repository"
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
