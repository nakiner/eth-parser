package v1

import (
	"context"

	pb "github.com/nakiner/eth-parser/pkg/pb/eth_parser/v1"
)

func (s *Service) GetCurrentBlock(ctx context.Context, request *pb.GetCurrentBlockRequest) (*pb.GetCurrentBlockResponse, error) {
	return &pb.GetCurrentBlockResponse{
		Block: s.ethService.GetBlockNumber(),
	}, nil
}
