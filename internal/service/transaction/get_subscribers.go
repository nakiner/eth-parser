package transaction

func (s *Service) GetSubscribers() []string {
	return s.subscriberRepo.GetSubscribers()
}
