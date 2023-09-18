package transaction

func (s *Service) SetBlockNumber(block int64) {
	s.blockRepo.SetBlockNumber(block)
}
