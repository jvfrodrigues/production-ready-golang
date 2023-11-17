package mocks

import (
	"github.com/jvfrodrigues/production-ready-golang/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) Register(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Find(id string) (*entities.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Transaction), args.Error(1)
}
