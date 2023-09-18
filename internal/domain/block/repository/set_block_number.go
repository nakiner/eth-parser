package repository

func (s *Service) SetBlockNumber(block int64) {
	s.cache.SetCurrentBlock(block)
}
