package routes

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/jvfrodrigues/production-ready-golang/internal/application/api/controllers"
	"github.com/jvfrodrigues/production-ready-golang/internal/application/treasury"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/repository"
)

func (r *Routes) setupTransactionRoutes() {
	store := persistence.NewRedisCache("cache:6379", "test", 24*time.Hour)
	exchangeService := treasury.NewTreasuryExchange()
	transactionRepository := repository.NewTrasactionRepositoryDb(r.Database)
	transactionController := controllers.NewTransactionController(transactionRepository, exchangeService)
	transactionGroup := r.Router.Group("/transaction")
	{
		transactionGroup.POST("/", transactionController.RegisterTransaction)
		transactionGroup.GET("/exchange/:id", cache.CachePage(store, 24*time.Hour, transactionController.FindTransactionAndExchangeCurrency))
	}
}
