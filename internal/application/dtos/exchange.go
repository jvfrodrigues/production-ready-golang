package dtos

type TreasuryExchangeRecord struct {
	RecordDate   string `json:"record_date"`
	Country      string `json:"country"`
	Currency     string `json:"currency"`
	ExchangeRate string `json:"exchange_rate"`
}

type TreasuryExchangeResponseDto struct {
	Data []TreasuryExchangeRecord `json:"data"`
}
