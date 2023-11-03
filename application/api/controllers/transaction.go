package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/application/usecases"
	"github.com/jvfrodrigues/transaction-product-wex/domain"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
	"github.com/jvfrodrigues/transaction-product-wex/infra/validator"
)

type TransactionController struct {
	RegisterTransactionUsecase usecases.RegisterTransactionUseCase
	ExchangeTransactionUsecase usecases.ExchangeTransactionUseCase
}

func NewTransactionController(transactionRepository entities.TransactionRepository, exchangeService domain.ExchangeService) *TransactionController {
	exchangeTransactionUsecase := usecases.NewExchangeTransactionUseCase(transactionRepository, exchangeService)
	registerTransactionUsecase := usecases.NewRegisterTransactionUseCase(transactionRepository)
	return &TransactionController{
		RegisterTransactionUsecase: *registerTransactionUsecase,
		ExchangeTransactionUsecase: *exchangeTransactionUsecase,
	}
}

func (txc *TransactionController) RegisterTransaction(ctx *gin.Context) {
	var request dtos.TransactionInputDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("test")
	err := validator.ValidateStruct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := txc.RegisterTransactionUsecase.Execute(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (txc *TransactionController) FindTransactionAndExchangeCurrency(ctx *gin.Context) {
	id := ctx.Param("id")
	country := ctx.Query("country")
	response, err := txc.ExchangeTransactionUsecase.Execute(id, country)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
