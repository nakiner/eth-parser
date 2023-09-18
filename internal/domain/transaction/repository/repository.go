package repository

import (
	"github.com/nakiner/eth-parser/internal/models"
)

type Service struct {
	cache transactionsCacheProvider
}

type transactionsCacheProvider interface {
	GetTransactions(address string) []models.Transaction
	AddTransactions(address string, transactions []models.Transaction)
}

func NewRepository(cache transactionsCacheProvider) *Service {
	return &Service{
		cache: cache,
	}
}
