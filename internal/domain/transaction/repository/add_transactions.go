package repository

import "github.com/nakiner/eth-parser/internal/models"

func (s *Service) AddTransactions(address string, transactions []models.Transaction) {
	s.cache.AddTransactions(address, transactions)
}
