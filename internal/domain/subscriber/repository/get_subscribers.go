package repository

func (s *Service) GetSubscribers() []string {
	return s.cache.GetSubscribers()
}
