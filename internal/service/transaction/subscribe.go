package transaction

func (s *Service) Subscribe(address string) bool {
	s.subscriberRepo.AddSubscriber(address)
	return true
}
