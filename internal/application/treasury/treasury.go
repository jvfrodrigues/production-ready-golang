package treasury

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/jvfrodrigues/production-ready-golang/internal/application/dtos"
	"github.com/jvfrodrigues/production-ready-golang/internal/domain"
	"github.com/jvfrodrigues/production-ready-golang/internal/infra/http"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TreasuryExchange struct {
	baseUrl string
}

func NewTreasuryExchange() *TreasuryExchange {
	return &TreasuryExchange{
		baseUrl: "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange",
	}
}

func (te *TreasuryExchange) GetCountryExchange(country string, transactionDate time.Time) ([]domain.CountryExchange, error) {
	var data dtos.TreasuryExchangeResponseDto
	formattedLimitDate := transactionDate.AddDate(0, -6, 0).Format("2006-01-02")
	formattedDate := transactionDate.Format("2006-01-02")
	formatTextCaser := cases.Title(language.English)
	formattedCountry := formatTextCaser.String(country)
	filter := fmt.Sprintf("country:in:(%s),record_date:lte:%s,record_date:gte:%s", formattedCountry, formattedDate, formattedLimitDate)
	query := url.Values{}
	query.Set("filter", filter)
	query.Set("sort", "-record_date")
	query.Set("page[number]", "1")
	query.Set("page[size]", "1")
	requestUrl := te.baseUrl + "?" + query.Encode()
	fmt.Println(requestUrl)
	response, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}
	var exchange []domain.CountryExchange
	if len(data.Data) < 1 {
		return exchange, nil
	}
	exchange = append(exchange, domain.CountryExchange{
		Country:      data.Data[0].Country,
		Currency:     data.Data[0].Currency,
		ExchangeRate: data.Data[0].ExchangeRate,
	})
	return exchange, nil
}
