package etherium

import (
	"context"

	"github.com/nakiner/eth-parser/internal/client/etherium"
	"github.com/nakiner/eth-parser/internal/models"
	"github.com/pkg/errors"
)

func (p *Provider) GetLogs(ctx context.Context, fromBlock int64, toBlock int64, address string) ([]models.Transaction, error) {
	resp, err := p.ethClient.GetLogs(ctx, fromBlock, toBlock, address)
	if err != nil {
		return nil, errors.Wrap(err, "provider.GetLogs error")
	}

	return transactionsResponseToModel(resp.Result), nil
}

func transactionsResponseToModel(data []etherium.ResultRow) []models.Transaction {
	res := make([]models.Transaction, len(data))
	for i, row := range data {
		res[i] = models.Transaction(row)
	}

	return res
}
