package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jvfrodrigues/transaction-product-wex/application/api/controllers"
	"github.com/jvfrodrigues/transaction-product-wex/application/dtos"
	"github.com/jvfrodrigues/transaction-product-wex/domain"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
	"github.com/jvfrodrigues/transaction-product-wex/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTransactionControllerRegisterTransaction(t *testing.T) {
	validTransaction, _ := entities.NewTransaction("test", time.Now(), 28)
	mockExchangeService := new(mocks.ExchangeServiceMock)
	testCases := []struct {
		testDescription      string
		expectedInput        interface{}
		shouldError          bool
		expectedUseCaseError error
		expectedResult       interface{}
		expectedResponseCode int
	}{
		{"Should register a transaction",
			&dtos.TransactionInputDto{
				Description:     "test",
				TransactionDate: time.Now(),
				Amount:          28,
			},
			false,
			nil,
			validTransaction,
			http.StatusCreated,
		},
		{"Should return error if usecase fails",
			&dtos.TransactionInputDto{
				Description:     "test",
				TransactionDate: time.Now(),
				Amount:          28,
			},
			true,
			errors.New("error on usecase"),
			gin.H{"error": "error on usecase"},
			http.StatusInternalServerError,
		},
		{"Should return error if input is not of the expected type",
			"invalid",
			true,
			nil,
			gin.H{"error": "json: cannot unmarshal string into Go value of type dtos.TransactionInputDto"},
			http.StatusBadRequest,
		},
		{"Should return error if input has any invalid data",
			&dtos.TransactionInputDto{
				Description:     "",
				TransactionDate: time.Now(),
				Amount:          28,
			},
			true,
			nil,
			gin.H{"error": "description is required"},
			http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			mockRepo := new(mocks.TransactionRepositoryMock)
			mockRepo.On("Register", mock.Anything).Return(tc.expectedUseCaseError)
			controller := controllers.NewTransactionController(mockRepo, mockExchangeService)
			expectedInputJSON, err := json.Marshal(tc.expectedInput)
			if err != nil {
				fmt.Println("Error marshalling the struct to JSON:", err)
				return
			}
			req, _ := http.NewRequest("POST", "/transaction", bytes.NewReader(expectedInputJSON))
			resp := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(resp)
			context.Request = req
			controller.RegisterTransaction(context)
			expectedResultJSON, err := json.Marshal(tc.expectedResult)
			if err != nil {
				fmt.Println("Error marshalling the struct to JSON:", err)
				return
			}
			assert.Equal(t, tc.expectedResponseCode, resp.Code)
			if tc.shouldError {
				assert.Equal(t, string(expectedResultJSON), resp.Body.String())
			} else {
				var responseStruct entities.Transaction
				err := json.Unmarshal(resp.Body.Bytes(), &responseStruct)
				if err != nil {
					fmt.Println("Error marshalling the struct to JSON:", err)
					return
				}
				require.Equal(t, tc.expectedResult.(*entities.Transaction).Description, responseStruct.Description)
				require.Equal(t, tc.expectedResult.(*entities.Transaction).PurchaseAmount, responseStruct.PurchaseAmount)
			}
		})
	}
}

func TestTransactionControllerFindTransactionAndExchangeCurrency(t *testing.T) {
	validTransaction, _ := entities.NewTransaction("test", time.Now(), 28)
	mockExchangeService := new(mocks.ExchangeServiceMock)
	mockExchangeService.On("GetCountryExchange", mock.Anything, mock.Anything).Return([]domain.CountryExchange{
		{
			Country:      "Brazil",
			Currency:     "Real",
			ExchangeRate: "5.033",
		},
	}, nil)
	testCases := []struct {
		testDescription      string
		expectedUseCaseError error
		expectedResult       interface{}
		expectedResponseCode int
	}{
		{"Should find transaction and return result",
			nil,
			&dtos.TransactionExchangedOutputDto{
				ID:              validTransaction.ID,
				Description:     validTransaction.Description,
				TransactionDate: validTransaction.TransactionDate,
				ExchangeRate:    "5.033",
				OriginalAmount:  "28.00",
				ConvertedAmount: "140.92",
			},
			http.StatusCreated,
		},
		{"Should return error if usecase fails",
			errors.New("error on usecase"),
			gin.H{"error": "error on usecase"},
			http.StatusInternalServerError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			mockRepo := new(mocks.TransactionRepositoryMock)
			mockRepo.On("Find", mock.Anything).Return(validTransaction, tc.expectedUseCaseError)
			controller := controllers.NewTransactionController(mockRepo, mockExchangeService)
			req, _ := http.NewRequest("GET", "/transaction/exchange/id?country=Brazil", nil)
			resp := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(resp)
			context.Request = req
			controller.FindTransactionAndExchangeCurrency(context)
			expectedResultJSON, err := json.Marshal(tc.expectedResult)
			if err != nil {
				fmt.Println("Error marshalling the struct to JSON:", err)
				return
			}
			assert.Equal(t, tc.expectedResponseCode, resp.Code)
			assert.Equal(t, string(expectedResultJSON), resp.Body.String())
		})
	}
}
