package transaction

func (s *Service) GetBlockNumber() int64 {
	return s.blockRepo.GetBlockNumber()
}
