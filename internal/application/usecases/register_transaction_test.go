package usecases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jvfrodrigues/production-ready-golang/internal/application/dtos"
	"github.com/jvfrodrigues/production-ready-golang/internal/application/usecases"
	"github.com/jvfrodrigues/production-ready-golang/internal/domain/entities"
	"github.com/jvfrodrigues/production-ready-golang/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterTransactionUseCase(t *testing.T) {
	validInput := dtos.TransactionInputDto{
		Description:     "Test",
		TransactionDate: time.Now(),
		Amount:          10.50,
	}
	validTransaction, _ := entities.NewTransaction(validInput.Description, validInput.TransactionDate, validInput.Amount)
	testCases := []struct {
		testDescription          string
		providedRequest          dtos.TransactionInputDto
		expectedTransactionError error
		expectedError            error
		expectedResult           *entities.Transaction
	}{
		{"Should register valid transaction",
			validInput,
			nil,
			nil,
			validTransaction,
		},
		{"Should give error on invalid input",
			dtos.TransactionInputDto{
				Description:     "",
				TransactionDate: time.Now(),
				Amount:          10.50,
			},
			nil,
			errors.New("description is required"),
			nil,
		},
		{"Should give error if repository fails to register",
			dtos.TransactionInputDto{
				Description:     "test",
				TransactionDate: time.Now(),
				Amount:          10.50,
			},
			errors.New("could not save on repository"),
			errors.New("could not save on repository"),
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			mockRepo := new(mocks.TransactionRepositoryMock)
			mockRepo.On("Register", mock.Anything).Return(tc.expectedTransactionError)
			txUseCase := usecases.NewRegisterTransactionUseCase(mockRepo)
			result, err := txUseCase.Execute(tc.providedRequest)
			if tc.expectedError != nil {
				assert.Error(t, err)
				require.Nil(t, result)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tc.expectedResult.Description, result.Description)
				require.Equal(t, tc.expectedResult.TransactionDate, result.TransactionDate)
				require.Equal(t, tc.expectedResult.PurchaseAmount, result.PurchaseAmount)
			}
		})
	}
}
