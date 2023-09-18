package transaction

import "github.com/nakiner/eth-parser/internal/models"

type Service struct {
	blockRepo       blockRepo
	subscriberRepo  subscriberRepo
	transactionRepo transactionRepo
}

type blockRepo interface {
	GetBlockNumber() int64
	SetBlockNumber(block int64)
}

type subscriberRepo interface {
	GetSubscribers() []string
	AddSubscriber(address string)
}

type transactionRepo interface {
	GetTransactions(address string) []models.Transaction
	AddTransactions(address string, transactions []models.Transaction)
}

func NewService(blockRepo blockRepo, subscriberRepo subscriberRepo, transactionRepo transactionRepo) *Service {
	return &Service{
		blockRepo:       blockRepo,
		subscriberRepo:  subscriberRepo,
		transactionRepo: transactionRepo,
	}
}
