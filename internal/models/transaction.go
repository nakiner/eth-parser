package models

import "github.com/nakiner/eth-parser/internal/client/etherium"

type Transaction struct {
	Address          string
	Topics           []string
	Data             string
	BlockNumber      string
	TransactionHash  string
	TransactionIndex string
	BlockHash        string
	LogIndex         string
	Removed          bool
}

func TransactionsResponseToModel(data []etherium.ResultRow) []Transaction {
	res := make([]Transaction, len(data))
	for i, row := range data {
		res[i] = Transaction(row)
	}

	return res
}
