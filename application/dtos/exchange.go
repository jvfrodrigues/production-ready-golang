package dtos

type ExchangeRecord struct {
	RecordDate   string `json:"record_date"`
	Country      string `json:"country"`
	Currency     string `json:"currency"`
	ExchangeRate string `json:"exchange_rate"`
}

type ExchangeResponseDto struct {
	Data []ExchangeRecord `json:"data"`
}
