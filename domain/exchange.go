package domain

import (
	"time"
)

type CountryExchange struct {
	Country      string `json:"country"`
	Currency     string `json:"currency"`
	ExchangeRate string `json:"exchange_rate"`
}

type ExchangeService interface {
	GetCountryExchange(country string, transactionDate time.Time) ([]CountryExchange, error)
}
