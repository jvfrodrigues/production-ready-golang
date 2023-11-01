package treasury

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/infra/http"
)

type TreasuryExchange struct {
	baseUrl string
}

func NewTreasuryExchange() *TreasuryExchange {
	return &TreasuryExchange{
		baseUrl: "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange",
	}
}

func (te *TreasuryExchange) GetCountryExchange(country string, transactionDate time.Time) (dtos.ExchangeResponseDto, error) {
	var data dtos.ExchangeResponseDto
	formattedLimitDate := transactionDate.AddDate(0, -6, 0).Format("2006-01-02")
	formattedDate := transactionDate.Format("2006-01-02")
	requestUrl := te.baseUrl + fmt.Sprintf("?filter=country:in:(%s),record_date:lte:%s,record_date:gte:%s&sort=-record_date&page[number]=1&page[size]=1", country, formattedDate, formattedLimitDate)
	response, err := http.Get(requestUrl)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
