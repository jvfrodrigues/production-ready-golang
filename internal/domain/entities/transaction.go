package entities

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	Register(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base            `valid:"required"`
	Description     string    `json:"description" gorm:"type:varchar(50)" valid:"required~description is required,stringlength(0|50)~description has a 50 character limit"`
	TransactionDate time.Time `json:"transaction_date" valid:"required"`
	PurchaseAmount  int64     `json:"purchase_amount" gorm:"bigint" valid:"required,int"`
}

func (transaction *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)
	if transaction.PurchaseAmount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	if err != nil {
		return err
	}
	return nil
}

func NewTransaction(description string, transactionDate time.Time, purchaseAmount float64) (*Transaction, error) {
	transaction := Transaction{
		Description:     description,
		TransactionDate: transactionDate,
		PurchaseAmount:  int64(purchaseAmount * 100),
	}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, err
}
