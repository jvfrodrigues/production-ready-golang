package domain

import (
	"time"
)

type ExchangeService interface {
	GetCountryExchange(country string, transactionDate time.Time) (interface{}, error)
}
