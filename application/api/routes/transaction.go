package routes

import (
	"github.com/jvfrodrigues/transaction-product-wex/application/api/controllers"
	"github.com/jvfrodrigues/transaction-product-wex/application/usecases"
	"github.com/jvfrodrigues/transaction-product-wex/infra/repository"
)

func (r *Routes) setupTransactionRoutes() {
	transactionRepository := repository.TransactionRepositoryDb{Db: r.Database}
	transactionUsecase := usecases.TransactionUseCase{TransactionRepository: transactionRepository}
	transactionController := controllers.NewTransactionController(transactionUsecase)
	transactionGroup := r.Router.Group("/transaction")
	{
		transactionGroup.POST("/", transactionController.RegisterTransaction)
		transactionGroup.GET("/exchange/:id/:country", transactionController.FindTransactionAndExchangeCurrency)
	}
}
