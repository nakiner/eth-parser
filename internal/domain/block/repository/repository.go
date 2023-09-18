package repository

type Service struct {
	cache blockNumberCacheProvider
}

type blockNumberCacheProvider interface {
	GetCurrentBlock() int64
	SetCurrentBlock(block int64)
}

func NewRepository(cache blockNumberCacheProvider) *Service {
	return &Service{
		cache: cache,
	}
}
