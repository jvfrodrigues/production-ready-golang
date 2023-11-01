package domain

import (
	"time"

	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
)

type ExchangeService interface {
	GetCountryExchange(country string, transactionDate time.Time) (dtos.ExchangeResponseDto, error)
}
