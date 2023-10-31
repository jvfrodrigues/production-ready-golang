package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/jvfrodrigues/transaction-product-wex/domain/entities"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r TransactionRepositoryDb) Register(transaction *entities.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepositoryDb) Find(id string) (*entities.Transaction, error) {
	var transaction entities.Transaction
	err := r.Db.First(&transaction, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
