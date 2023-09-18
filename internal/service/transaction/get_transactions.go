package transaction

import "github.com/nakiner/eth-parser/internal/models"

func (s *Service) GetTransactions(address string) []models.Transaction {
	return s.transactionRepo.GetTransactions(address)
}
