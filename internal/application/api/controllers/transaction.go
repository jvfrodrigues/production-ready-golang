package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jvfrodrigues/production-ready-golang/internal/application/dtos"
	"github.com/jvfrodrigues/production-ready-golang/internal/application/usecases"
	"github.com/jvfrodrigues/production-ready-golang/internal/domain"
	"github.com/jvfrodrigues/production-ready-golang/internal/domain/entities"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/logger"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/logger/zap"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/validator"
)

type TransactionController struct {
	RegisterTransactionUsecase usecases.RegisterTransactionUseCase
	ExchangeTransactionUsecase usecases.ExchangeTransactionUseCase
	Logger                     logger.ILogger
}

func NewTransactionController(transactionRepository entities.TransactionRepository, exchangeService domain.ExchangeService) *TransactionController {
	exchangeTransactionUsecase := usecases.NewExchangeTransactionUseCase(transactionRepository, exchangeService)
	registerTransactionUsecase := usecases.NewRegisterTransactionUseCase(transactionRepository)
	logger := zap.NewLogger()
	return &TransactionController{
		RegisterTransactionUsecase: *registerTransactionUsecase,
		ExchangeTransactionUsecase: *exchangeTransactionUsecase,
		Logger:                     logger,
	}
}

func (txc *TransactionController) RegisterTransaction(ctx *gin.Context) {
	txc.Logger.Info("Executing Register Transaction Usecase")
	var request dtos.TransactionInputDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		txc.Logger.Error(fmt.Sprintf("Error on getting request body %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("test")
	err := validator.ValidateStruct(request)
	if err != nil {
		txc.Logger.Error(fmt.Sprintf("Error on request body validation %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := txc.RegisterTransactionUsecase.Execute(request)
	if err != nil {
		txc.Logger.Error(fmt.Sprintf("Error on Register Transaction Usecase %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (txc *TransactionController) FindTransactionAndExchangeCurrency(ctx *gin.Context) {
	txc.Logger.Info("Executing Find Transaction and Exchange Currency Usecase")
	id := ctx.Param("id")
	country := ctx.Query("country")
	response, err := txc.ExchangeTransactionUsecase.Execute(id, country)
	if err != nil {
		txc.Logger.Error(fmt.Sprintf("Error on Exchange Transaction Usecase %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
