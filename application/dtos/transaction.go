package dtos

import "time"

type TransactionInputDto struct {
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date"`
	Amount          float64   `json:"amount"`
}

type TransactionExchangedOutputDto struct {
	ID              string    `json:"id"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date"`
	ExchangeRate    float64   `json:"exchange rate"`
	OriginalAmount  string    `json:"original_amount"`
	ConvertedAmount string    `json:"converted_amount"`
}
