package v1

import (
	"context"
	"strings"

	pb "github.com/nakiner/eth-parser/pkg/pb/eth_parser/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Subscribe(ctx context.Context, request *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	addr := request.GetAddress()
	if len(addr) < 1 {
		return nil, status.Error(codes.InvalidArgument, "address should be > 0")
	}
	return &pb.SubscribeResponse{
		Status: s.ethService.Subscribe(strings.ToLower(addr)),
	}, nil
}
