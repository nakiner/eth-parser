package transaction

import "github.com/nakiner/eth-parser/internal/models"

func (s *Service) AddTransactions(address string, transactions []models.Transaction) {
	s.transactionRepo.AddTransactions(address, transactions)
}
