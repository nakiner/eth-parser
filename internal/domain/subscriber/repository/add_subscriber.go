package repository

func (s *Service) AddSubscriber(address string) {
	s.cache.AddSubscriber(address)
}
