package usecases

import (
	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
)

type RegisterTransactionUseCase struct {
	TransactionRepository entities.TransactionRepository
}

func NewRegisterTransactionUseCase(transactionRepository entities.TransactionRepository) *RegisterTransactionUseCase {
	return &RegisterTransactionUseCase{
		TransactionRepository: transactionRepository,
	}
}

func (tx RegisterTransactionUseCase) Execute(request dtos.TransactionInputDto) (*entities.Transaction, error) {
	transaction, err := entities.NewTransaction(request.Description, request.TransactionDate, request.Amount)
	if err != nil {
		return nil, err
	}
	err = tx.TransactionRepository.Register(transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
