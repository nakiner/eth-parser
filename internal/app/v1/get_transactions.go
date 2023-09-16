package v1

import (
	"context"

	pb "github.com/nakiner/eth-parser/pkg/pb/eth_parser/v1"
)

func (s Service) GetTransactions(ctx context.Context, request *pb.GetTransactionsRequest) (*pb.GetTransactionsResponse, error) {
	return &pb.GetTransactionsResponse{}, nil
}
