package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/application/usecases"
)

type TransactionController struct {
	TransactionUsecase usecases.TransactionUseCase
}

func NewTransactionController(transactionUsecase usecases.TransactionUseCase) *TransactionController {
	return &TransactionController{
		TransactionUsecase: transactionUsecase,
	}
}

func (txc *TransactionController) RegisterTransaction(ctx *gin.Context) {
	var request dtos.TransactionInputDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	response, err := txc.TransactionUsecase.Register(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (txc *TransactionController) FindTransactionAndExchangeCurrency(ctx *gin.Context) {
	id := ctx.Param("id")
	country := ctx.Param("country")
	response, err := txc.TransactionUsecase.FindAndExchangeCurrency(id, country)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
