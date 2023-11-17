package dtos

import "time"

type TransactionInputDto struct {
	Description     string    `json:"description" valid:"required~description is required"`
	TransactionDate time.Time `json:"transaction_date" valid:"required~transaction date is required"`
	Amount          float64   `json:"amount" valid:"required,float~amount must be a number"`
}

type TransactionExchangedOutputDto struct {
	ID              string    `json:"id"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date"`
	ExchangeRate    string    `json:"exchange rate"`
	OriginalAmount  string    `json:"original_amount"`
	ConvertedAmount string    `json:"converted_amount"`
}
