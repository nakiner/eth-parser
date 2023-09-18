package repository

import "github.com/nakiner/eth-parser/internal/models"

func (s *Service) GetTransactions(address string) []models.Transaction {
	return s.cache.GetTransactions(address)
}
