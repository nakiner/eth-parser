package v1

import (
	"context"

	pb "github.com/nakiner/eth-parser/pkg/pb/eth_parser/v1"
)

func (s Service) Subscribe(ctx context.Context, request *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	return &pb.SubscribeResponse{}, nil
}
