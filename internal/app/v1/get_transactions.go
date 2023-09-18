package v1

import (
	"context"
	"strings"

	"github.com/nakiner/eth-parser/internal/models"
	pb "github.com/nakiner/eth-parser/pkg/pb/eth_parser/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetTransactions(ctx context.Context, request *pb.GetTransactionsRequest) (*pb.GetTransactionsResponse, error) {
	address := request.GetAddress()
	if len(address) < 1 {
		return nil, status.Error(codes.InvalidArgument, "address cannot be empty")
	}
	transactions := s.ethService.GetTransactions(strings.ToLower(address))
	return &pb.GetTransactionsResponse{
		Transactions: convertTransactionsToPB(transactions),
	}, nil
}

func convertTransactionsToPB(data []models.Transaction) []*pb.Transaction {
	res := make([]*pb.Transaction, len(data))
	for i, row := range data {
		res[i] = &pb.Transaction{
			Address:          row.Address,
			Topics:           row.Topics,
			Data:             row.Data,
			BlockNumber:      row.BlockNumber,
			TransactionHash:  row.TransactionHash,
			TransactionIndex: row.TransactionIndex,
			BlockHash:        row.BlockHash,
			LogIndex:         row.LogIndex,
			Removed:          row.Removed,
		}
	}
	return res
}
