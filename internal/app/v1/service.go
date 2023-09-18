package v1

import "github.com/nakiner/eth-parser/internal/models"

type Service struct {
	ethService ethService
}

type ethService interface {
	GetBlockNumber() int64
	Subscribe(address string) bool
	GetTransactions(address string) []models.Transaction
}

func NewService(ethService ethService) *Service {
	return &Service{
		ethService: ethService,
	}
}
