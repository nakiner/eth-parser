package etherium

import (
	"context"

	"github.com/nakiner/eth-parser/internal/models"
	"github.com/pkg/errors"
)

func (p *Provider) GetLogs(ctx context.Context, fromBlock int64, toBlock int64, address string) ([]models.Transaction, error) {
	resp, err := p.ethClient.GetLogs(ctx, fromBlock, toBlock, address)
	if err != nil {
		return nil, errors.Wrap(err, "provider.GetLogs error")
	}

	return models.TransactionsResponseToModel(resp.Result), nil
}
