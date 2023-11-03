package usecases

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/domain"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
)

type ExchangeTransactionUseCase struct {
	TransactionRepository entities.TransactionRepository
	ExchangeService       domain.ExchangeService
}

func NewExchangeTransactionUseCase(transactionRepository entities.TransactionRepository, exchangeService domain.ExchangeService) *ExchangeTransactionUseCase {
	return &ExchangeTransactionUseCase{
		TransactionRepository: transactionRepository,
		ExchangeService:       exchangeService,
	}
}

func (tx ExchangeTransactionUseCase) Execute(id string, country string) (*dtos.TransactionExchangedOutputDto, error) {
	transaction, err := tx.TransactionRepository.Find(id)
	if err != nil {
		return nil, err
	}
	exchange, err := tx.ExchangeService.GetCountryExchange(country, transaction.TransactionDate)
	if err != nil {
		return nil, err
	}
	if len(exchange) <= 0 {
		return nil, errors.New("no exchange information found for country in 6 month period")
	}
	originalAmount := new(big.Rat).SetFrac64(transaction.PurchaseAmount, 100)
	parsedExchangeRate, err := strconv.ParseFloat(exchange[0].ExchangeRate, 64)
	if err != nil {
		return nil, err
	}
	exchangeRate := new(big.Rat).SetFloat64(parsedExchangeRate)
	convertedAmount := new(big.Rat).Mul(originalAmount, exchangeRate)
	response := &dtos.TransactionExchangedOutputDto{
		ID:              transaction.ID,
		Description:     transaction.Description,
		TransactionDate: transaction.TransactionDate,
		ExchangeRate:    exchange[0].ExchangeRate,
		OriginalAmount:  originalAmount.FloatString(2),
		ConvertedAmount: convertedAmount.FloatString(2),
	}
	return response, nil
}
