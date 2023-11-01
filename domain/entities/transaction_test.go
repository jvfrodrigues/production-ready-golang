package entities_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionCreation(t *testing.T) {
	testCases := []struct {
		testDescription string
		description     string
		transactionDate time.Time
		purchaseAmount  float64
		expectedError   error
	}{
		{"Valid data on fields should create transaction", "test", time.Now(), 10.55, nil},
		{"Description field empty should give an error", "", time.Now(), 10.55, errors.New("description is required")},
		{"Negative amount should give an error", "test", time.Now(), -10.55, errors.New("amount must be greater than 0")},
		{"Description longer than 50 characters should give an error", "test974812748-1242119-2048-90248118qwsiurioq19287408912478", time.Now(), 10.55, errors.New("description has a 50 character limit")},
	}

	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			transaction, err := entities.NewTransaction(tc.description, tc.transactionDate, tc.purchaseAmount)

			if tc.expectedError != nil {
				assert.Error(t, err)
				require.Nil(t, transaction)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.Nil(t, err)
				require.NotEmpty(t, uuid.FromStringOrNil(transaction.ID))
			}
		})
	}
}
