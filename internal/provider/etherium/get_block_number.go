package etherium

import "context"

func (p *Provider) GetBlockNumber(ctx context.Context) (int64, error) {
	return p.ethClient.GetBlockNumber(ctx)
}
