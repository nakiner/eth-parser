package etherium

import (
	"context"

	"github.com/nakiner/eth-parser/internal/client/etherium"
)

type Provider struct {
	ethClient etheriumClient
}

type etheriumClient interface {
	GetBlockNumber(ctx context.Context) (int64, error)
	GetLogs(ctx context.Context, fromBlock int64, toBlock int64, address string) (etherium.GetLogsResponse, error)
}

func NewProvider(ethClient etheriumClient) *Provider {
	return &Provider{
		ethClient: ethClient,
	}
}
