package repository

func (s *Service) GetBlockNumber() int64 {
	return s.cache.GetCurrentBlock()
}
