package usecases

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/application/exchange"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
)

type TransactionUseCase struct {
	TransactionRepository entities.TransactionRepository
}

func (tx TransactionUseCase) Register(request dtos.TransactionInputDto) (*entities.Transaction, error) {
	transaction, err := entities.NewTransaction(request.Description, request.TransactionDate, request.Amount)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", transaction)
	err = tx.TransactionRepository.Register(transaction)
	if transaction.ID == "" {
		return nil, err
	}
	return transaction, nil
}

func (tx TransactionUseCase) FindAndExchangeCurrency(id string, country string) (*dtos.TransactionExchangedOutputDto, error) {
	transaction, err := tx.TransactionRepository.Find(id)
	if err != nil {
		return nil, err
	}
	exchange, err := exchange.GetCountryExchange(country, transaction.TransactionDate)
	if err != nil {
		return nil, err
	}
	if len(exchange.Data) <= 0 {
		return nil, errors.New("no exchange information found for country in 6 month period")
	}
	originalAmount := new(big.Rat).SetFrac64(transaction.PurchaseAmount, 100)
	parsedExchangeRate, err := strconv.ParseFloat(exchange.Data[0].ExchangeRate, 64)
	if err != nil {
		return nil, err
	}
	exchangeRate := new(big.Rat).SetFloat64(parsedExchangeRate)
	convertedAmount := new(big.Rat).Mul(originalAmount, exchangeRate)
	response := &dtos.TransactionExchangedOutputDto{
		ID:              transaction.ID,
		Description:     transaction.Description,
		TransactionDate: transaction.TransactionDate,
		ExchangeRate:    parsedExchangeRate,
		OriginalAmount:  originalAmount.FloatString(2),
		ConvertedAmount: convertedAmount.FloatString(2),
	}
	return response, nil
}
