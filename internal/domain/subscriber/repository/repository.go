package repository

type Service struct {
	cache subscriberCacheProvider
}

type subscriberCacheProvider interface {
	AddSubscriber(address string)
	GetSubscribers() []string
}

func NewRepository(cache subscriberCacheProvider) *Service {
	return &Service{
		cache: cache,
	}
}
