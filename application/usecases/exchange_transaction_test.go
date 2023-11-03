package usecases_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/application/usecases"
	"github.com/jvfrodrigues/transaction-product-wex/domain"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
	"github.com/jvfrodrigues/transaction-product-wex/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExchangeTransactionUseCase(t *testing.T) {
	validTransaction, _ := entities.NewTransaction("test", time.Now(), 28)
	testCases := []struct {
		testDescription         string
		foundTransaction        *entities.Transaction
		foundExchange           []domain.CountryExchange
		country                 string
		expectedRepositoryError error
		expectedServiceError    error
		expectedError           error
		expectedResult          *dtos.TransactionExchangedOutputDto
	}{
		{"Should find valid transaction and get currency requested",
			validTransaction,
			[]domain.CountryExchange{
				{
					Country:      "Brazil",
					Currency:     "Real",
					ExchangeRate: "5.033",
				},
			},
			"Brazil",
			nil,
			nil,
			nil,
			&dtos.TransactionExchangedOutputDto{
				ID:              validTransaction.ID,
				Description:     validTransaction.Description,
				TransactionDate: validTransaction.TransactionDate,
				ExchangeRate:    "5.033",
				OriginalAmount:  "28.00",
				ConvertedAmount: "140.92",
			},
		},
		{"Should not find transaction and give error",
			validTransaction,
			[]domain.CountryExchange{
				{
					Country:      "Brazil",
					Currency:     "Real",
					ExchangeRate: "5.033",
				},
			},
			"Brazil",
			errors.New("transaction not found"),
			nil,
			errors.New("transaction not found"),
			nil,
		},
		{"Should not find exchange information in period and give error",
			validTransaction,
			[]domain.CountryExchange{},
			"Brazil",
			nil,
			nil,
			errors.New("no exchange information found for country in 6 month period"),
			nil,
		},
		{"Should exchange service give error return error",
			validTransaction,
			[]domain.CountryExchange{},
			"Brazil",
			nil,
			errors.New("request error 500"),
			errors.New("request error 500"),
			nil,
		},
		{"Should find invalid exchange rate and return error",
			validTransaction,
			[]domain.CountryExchange{
				{
					Country:      "Brazil",
					Currency:     "Real",
					ExchangeRate: "invalid",
				},
			},
			"Brazil",
			nil,
			nil,
			errors.New(`strconv.ParseFloat: parsing "invalid": invalid syntax`),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			mockRepo := new(mocks.TransactionRepositoryMock)
			mockExchangeService := new(mocks.ExchangeServiceMock)
			mockRepo.On("Find", tc.foundTransaction.ID).Return(tc.foundTransaction, tc.expectedRepositoryError)
			mockExchangeService.On("GetCountryExchange", tc.country, tc.foundTransaction.TransactionDate).Return(tc.foundExchange, tc.expectedServiceError)
			txUseCase := usecases.NewExchangeTransactionUseCase(mockRepo, mockExchangeService)
			result, err := txUseCase.Execute(tc.foundTransaction.ID, tc.country)
			if tc.expectedError != nil {
				assert.Error(t, err)
				require.Nil(t, result)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.Nil(t, err)
				if !reflect.DeepEqual(result, tc.expectedResult) {
					t.Errorf("Expected response %v, but got %v", tc.expectedResult, result)
				}
			}
		})
	}
}
