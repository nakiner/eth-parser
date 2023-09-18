package cache

import "github.com/nakiner/eth-parser/internal/models"

func (p *Provider) GetTransactions(address string) []models.Transaction {
	res, ok := p.transactions.Get(address)
	if !ok {
		return nil
	}
	return res.([]models.Transaction)
}

func (p *Provider) AddTransactions(address string, transactions []models.Transaction) {
	data := p.GetTransactions(address)
	if len(data) < 1 {
		p.transactions.Set(address, transactions)
		return
	}

	data = append(data, transactions...)
	p.transactions.Set(address, data)
}
