package cache

type inmemoryStorage interface {
	Get(key string) (any, bool)
	Set(key string, val any)
}

type Provider struct {
	currentBlock inmemoryStorage
	subscribers  inmemoryStorage
	transactions inmemoryStorage
}

func NewCacheProvider(currentBlock inmemoryStorage, subscribers inmemoryStorage, transactions inmemoryStorage) *Provider {
	return &Provider{
		currentBlock: currentBlock,
		subscribers:  subscribers,
		transactions: transactions,
	}
}
