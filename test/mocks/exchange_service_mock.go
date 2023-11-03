package mocks

import (
	"time"

	"github.com/jvfrodrigues/transaction-product-wex/domain"
	"github.com/stretchr/testify/mock"
)

type ExchangeServiceMock struct {
	mock.Mock
}

func (m *ExchangeServiceMock) GetCountryExchange(country string, transactionDate time.Time) ([]domain.CountryExchange, error) {
	args := m.Called(country, transactionDate)
	return args.Get(0).([]domain.CountryExchange), args.Error(1)
}
