package models

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
